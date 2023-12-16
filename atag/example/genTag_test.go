package example

import (
	"fmt"
	"gitee.com/asktop_golib/util/atag"
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

func TestGenTag1(t *testing.T) {
	tag := atag.GenTag(`
	Id       int64
	UserName string //姓名
	IdCard   string //身份证
	`,
		"json", "form")
	fmt.Println(tag)
}

func TestGenTag2(t *testing.T) {
	tag := atag.GenTag(`
type demo struct {
	Id       int64
	UserName string //姓名
	IdCard   string //身份证
	user
}
	`,
		"json", "form")
	fmt.Println(tag)
}

func TestGenTag3(t *testing.T) {
	tag := atag.GenTag("type demo struct {\n\tId       int64\n\tUserName string //姓名\n\tIdCard   string `json:\"card_no\"` //身份证\n}",
		"json", "db")
	fmt.Println(tag)
}
