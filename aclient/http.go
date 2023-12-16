package aclient

import (
    "bytes"
    "crypto/tls"
    "crypto/x509"
    "encoding/json"
    "fmt"
    "gitee.com/asktop_golib/util/acast"
    "io"
    "io/ioutil"
    "net/http"
    "net/url"
    "strings"
    "time"
)

type Client struct {
    client *http.Client
}

func NewClient(certPath ...string) *Client {
    c := new(Client)
    tlsConfig := &tls.Config{
        InsecureSkipVerify: true,
    }
    if len(certPath) > 0 {
        //加载双向认证证书
        crt, err := ioutil.ReadFile(certPath[0])
        if err != nil {
            fmt.Println("获取客户端https证书出错，err:", err)
        } else {
            crts := x509.NewCertPool()
            crts.AppendCertsFromPEM(crt)
            tlsConfig.RootCAs = crts
        }
    }
    //创建HttpClient并发起请求
    c.client = &http.Client{
        Transport: &http.Transport{
            DisableKeepAlives:   true, //true:不同HTTP请求之间TCP连接的重用将被阻止（http1.1默认为长连接，此处改为短连接）
            MaxIdleConnsPerHost: 512,  //控制每个主机下的最大闲置连接数目
            TLSClientConfig:     tlsConfig,
            TLSHandshakeTimeout: time.Second * 10,
        },
        Timeout: time.Second * 10, //Client请求的时间限制,该超时限制包括连接时间、重定向和读取response body时间;Timeout为零值表示不设置超时
    }
    return c
}

func (c *Client) SetClient(client *http.Client) *Client {
    c.client = client
    return c
}

func (c *Client) SetTimeout(timeout time.Duration) *Client {
    c.client.Timeout = timeout
    return c
}

func (c *Client) Request(method string, url string, body io.Reader, headers ...map[string]string) (respBody []byte, statusCode int, err error) {
    //创建HttpRequest
    req, err := http.NewRequest(strings.ToUpper(method), url, body)
    if err != nil {
        return nil, 0, err
    }
    //req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    //req.Header.Set("Content-Type", "application/json")
    if len(headers) > 0 {
        for k, v := range headers[0] {
            req.Header.Set(k, v)
        }
    }

    resp, err := c.client.Do(req)
    if resp != nil {
        statusCode = resp.StatusCode
    }
    if err != nil {
        return nil, statusCode, err
    }
    defer resp.Body.Close()

    //解析响应信息
    reply, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, statusCode, err
    }
    return reply, statusCode, nil
}

func (c *Client) Get(URL string, querys map[string]interface{}, headers ...map[string]string) (respBody []byte, statusCode int, err error) {
    if len(querys) > 0 {
        var rawQuerys []string
        for key, value := range querys {
            rawQuerys = append(rawQuerys, fmt.Sprintf("%s=%v", key, value))
        }
        if strings.Contains(URL, "?") {
            URL = fmt.Sprintf("%s&%s", URL, strings.Join(rawQuerys, "&"))
        } else {
            URL = fmt.Sprintf("%s?%s", URL, strings.Join(rawQuerys, "&"))
        }
    }
    return c.Request(http.MethodGet, URL, nil, headers...)
}

func (c *Client) Post(URL string, params map[string]interface{}, headers ...map[string]string) (respBody []byte, statusCode int, err error) {
    data := url.Values{}
    for k, v := range params {
        var nv []string
        switch vtemp := v.(type) {
        case string:
            nv = append(nv, vtemp)
        default:
            var e error
            nv, e = acast.ToStringSliceE(v)
            if e != nil {
                return nil, 0, e
            }
        }
        data[k] = nv
    }
    body := strings.NewReader(data.Encode())
    return c.Request(http.MethodPost, URL, body, headers...)
}

func (c *Client) PostForm(URL string, params map[string]interface{}, headers ...map[string]string) (respBody []byte, statusCode int, err error) {
    header := make(map[string]string)
    if len(headers) > 0 && headers[0] != nil {
        header = headers[0]
    }
    if _, ok := header["Content-Type"]; !ok {
        header["Content-Type"] = "application/x-www-form-urlencoded"
    }
    return c.Post(URL, params, header)
}

func (c *Client) PostJson(URL string, params map[string]interface{}, headers ...map[string]string) (respBody []byte, statusCode int, err error) {
    header := make(map[string]string)
    if len(headers) > 0 && headers[0] != nil {
        header = headers[0]
    }
    if _, ok := header["Content-Type"]; !ok {
        header["Content-Type"] = "application/json"
    }
    //组装请求信息
    data, _ := json.Marshal(params)
    body := bytes.NewReader(data)
    return c.Request(http.MethodPost, URL, body, header)
}
