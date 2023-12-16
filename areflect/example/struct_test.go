package example

import (
	"fmt"
	"gitee.com/asktop_golib/util/areflect"
	"testing"
)

type demo struct {
	Id       int64
	UserName string //姓名
	IdCard   string `json:"card_no"` //身份证
	user
}

type user struct {
	Truename string `json:"truename"`
	Sex      int64
}

func TestGetStructFieldNames(t *testing.T) {
	fmt.Println(areflect.GetStructFieldNames(new(demo), ""))
	fmt.Println(areflect.GetStructFieldNamesWithoutTag(new(demo), "", "id"))
	fmt.Println(areflect.GetStructFieldNames(new(demo), "json"))
}
