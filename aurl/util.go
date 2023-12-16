package aurl

import (
    "fmt"
    "net/url"
    "sort"
    "strings"
)

//获取url的网址
func GetUrlSite(urlstr string) string {
    if !strings.HasPrefix(urlstr, "http://") && !strings.HasPrefix(urlstr, "https://") {
        return ""
    }
    urlObj, err := url.Parse(urlstr)
    if err != nil {
        return ""
    }
    site := fmt.Sprintf("%s://%s", urlObj.Scheme, urlObj.Host)
    return site
}

func GetUrlHost(urlstr string) string {
    if !strings.HasPrefix(urlstr, "http://") && !strings.HasPrefix(urlstr, "https://") {
        return ""
    }
    urlObj, err := url.Parse(urlstr)
    if err != nil {
        return ""
    }
    return urlObj.Host
}

//拼接URL
func JoinURL(uris ...string) string {
    var url string
    var urisTemp []string
    for _, uri := range uris {
        uri = strings.Trim(uri, " ")
        uri = strings.Trim(uri, "/")
        uri = strings.Trim(uri, `\`)
        if uri != "" {
            urisTemp = append(urisTemp, uri)
        }
    }
    url = strings.Join(urisTemp, "/")
    if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
        url = "/" + url
    }
    return url
}

//排序 拼接
func SortParams(m map[string]string) string {
    //排序
    keys := make([]string, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    sort.Strings(keys)

    //拼接
    var buf strings.Builder
    for _, k := range keys {
        if buf.Len() > 0 {
            buf.WriteByte('&')
        }
        buf.WriteString(k)
        buf.WriteByte('=')
        buf.WriteString(m[k])
    }
    return buf.String()
}

//排序 编码 拼接
func SortEncodeParams(params map[string]string) string {
    values := url.Values{}
    for k, v := range params {
        values.Add(k, v)
    }
    return values.Encode()
}
