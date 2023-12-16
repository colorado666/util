package aguid

import (
    "encoding/json"
    "gitee.com/asktop_golib/util/acast"
    "gitee.com/asktop_golib/util/afile"
    "github.com/wenzhenxi/gorsa"
    "strings"
)

//用户信息
type guidInfo struct {
    Iat  int64                  `json:"iat"`
    Exp  int64                  `json:"exp"`
    Data map[string]interface{} `json:"data"`
}

//Rsa加密
func RsaEncrypt(data map[string]interface{}, exp int64, privateKeyStr string) (guid string, err error) {
    info := new(guidInfo)
    info.Exp = exp
    info.Data = data
    body, err := json.Marshal(info)
    if err != nil {
        return "", err
    }

    //rsa加密
    guid, err = gorsa.PriKeyEncrypt(string(body), privateKeyStr)
    if err != nil {
        return "", err
    }

    //自定义处理
    if strings.Contains(guid, "+") || strings.Contains(guid, "/") {
        guid = strings.Replace(guid, "+", "-", -1)
        guid = strings.Replace(guid, "/", "_", -1)
        guid = strings.TrimSuffix(guid, "=")
    }
    return guid, nil
}

//Rsa加密
func RsaEncryptObj(data interface{}, exp int64, privateKeyStr string) (guid string, err error) {
    dataMap := map[string]interface{}{}
    err = acast.StructToMap(data, &dataMap)
    if err != nil {
        return "", err
    }
    return RsaEncrypt(dataMap, exp, privateKeyStr)
}

//Rsa加密
func RsaPathEncrypt(data map[string]interface{}, exp int64, privateKeyPath string) (guid string, err error) {
    privateKeyStr, err := afile.ReadFile(privateKeyPath)
    if err != nil {
        return "", err
    }
    return RsaEncrypt(data, exp, privateKeyStr)
}

//Rsa加密
func RsaPathEncryptObj(data interface{}, exp int64, privateKeyPath string) (guid string, err error) {
    privateKeyStr, err := afile.ReadFile(privateKeyPath)
    if err != nil {
        return "", err
    }
    return RsaEncryptObj(data, exp, privateKeyStr)
}

//Rsa解密
func RsaDecrypt(guid string, publicKeyStr string) (data map[string]interface{}, err error) {
    //自定义处理
    if strings.Contains(guid, "-") || strings.Contains(guid, "_") {
        //补等号
        remainder := len(guid) % 4
        if remainder != 0 {
            //该补的等号的数量
            padlen := 4 - remainder
            for i := 0; i < padlen; i++ {
                guid = guid + "="
            }
        }
        guid = strings.Replace(guid, "-", "+", -1)
        guid = strings.Replace(guid, "_", "/", -1)
    }

    //rsa解密
    body, err := gorsa.PublicDecrypt(guid, publicKeyStr)
    if err != nil {
        return nil, err
    }

    var info guidInfo
    err = json.Unmarshal([]byte(body), &info)
    if err != nil {
        return nil, err
    }
    //if info.Exp > 0 && info.Exp < time.Now().Unix() {
    //    return nil, errors.New("guid is expired")
    //}

    return info.Data, nil
}

//Rsa解密
func RsaDecryptObj(guid string, publicKeyStr string, data interface{}) error {
    dataMap, err := RsaDecrypt(guid, publicKeyStr)
    if err != nil {
        return err
    }
    return acast.MapToStruct(dataMap, data)
}

//Rsa解密
func RsaPathDecrypt(guid string, publicKeyPath string) (data map[string]interface{}, err error) {
    publicKeyStr, err := afile.ReadFile(publicKeyPath)
    if err != nil {
        return nil, err
    }
    return RsaDecrypt(guid, publicKeyStr)
}

//Rsa解密
func RsaPathDecryptObj(guid string, publicKeyPath string, data interface{}) error {
    publicKeyStr, err := afile.ReadFile(publicKeyPath)
    if err != nil {
        return err
    }
    return RsaDecryptObj(guid, publicKeyStr, data)
}
