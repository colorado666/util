package aguid

import (
    "fmt"
    "testing"
    "time"
)

var (
    privateKey = `
-----BEGIN PRIVATE KEY-----
MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAKmOMrDNDEMgapVC
HM+0hK36OPWVCzw4lRtPtxOSobV9C5h5y+RGSEbRwoZvHXntl19iuTwBMu9PXaNA
sE+WFvkpSRy0K0PcSsQ0p+hnqbumHvz6eAxnXLW1jZBi/3HE7xcZWDQZHk03jPVk
s2lMdsu2PG5TVUtAocq2ThEUDT+LAgMBAAECgYAQS+2NVrf1/7iezfLs98HE4wb0
e7XPvPR/4oKLLA3E3tbtec9iCmtJ+0FCII3puS9SaK+7F7Zoj+1FEqOfkqJ0eh+Z
Hbs3kSGrLTnUqhx6hfbbZ7KmOb5ER9YoXjMLiJa6GjrI3JSvTTX4Gq4gtVOmmgKo
Sk1SGw4hDeIRcjqf0QJBAN++0JIbsPfnhp9LC8oeyyh4sJlnXNRdgg35VF2pntLG
4w8MM+tulXz+hoh/yjipiynq+f5FMrUILxLj+Yg7M7kCQQDB/4m7Yx7OZB9E7uQD
tZZKJDsf6jOaITe27HVkVoaYfdjiJdWxgTqOu+XfUlSamDfS8T577ozYe+UJQbjC
J7djAkAdV/Yco1sTOB8Utw/lwyIbvbBTfhXTmCvdT0y8N+mndQQETjJk9wzN5seY
EUJhKgMhTTraGRMsYNVjodhxAGqpAkEAp2hGiKfua5hGy4uCxBitCmndg5rs0LKd
i4LdgqqQF4Nx5yVTFIw55fYLVf2L+KQmeCeqDfQ7Io03hhdqwlDXlQJBAISJqBbK
a/CqrACq9IsQgUaN+OSrVouJWieEYwRYCSHuHFN6nF3q4mOOeQ/qs03NKEL2y5Ad
ZbT6qUiSgz/zHdY=
-----END PRIVATE KEY-----
`
    publicKey = `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCpjjKwzQxDIGqVQhzPtISt+jj1
lQs8OJUbT7cTkqG1fQuYecvkRkhG0cKGbx157ZdfYrk8ATLvT12jQLBPlhb5KUkc
tCtD3ErENKfoZ6m7ph78+ngMZ1y1tY2QYv9xxO8XGVg0GR5NN4z1ZLNpTHbLtjxu
U1VLQKHKtk4RFA0/iwIDAQAB
-----END PUBLIC KEY-----
`
)

func TestGuid(t *testing.T) {
    dataMap := map[string]interface{}{}
    dataMap["username"] = "888888"
    guid, err := RsaEncrypt(dataMap, time.Now().Add(time.Hour).Unix(), privateKey)
    if err != nil {
        fmt.Println(err)
        return
    } else {
        fmt.Println("guid:")
        fmt.Println(guid)
    }

    info, err := RsaDecrypt(guid, publicKey)
    if err != nil {
        fmt.Println(err)
        return
    } else {
        fmt.Println("info:")
        fmt.Println(info)
    }
}
