package aurl

import (
    "fmt"
    "testing"
)

func TestGetSite(t *testing.T) {
    site := GetUrlSite("https://v3master-oss.dsceshi.cn/uploads/20210919/f56932ac-d468-4be3-5a8c-1ae715b633d8.png")
    fmt.Println(site)
}
