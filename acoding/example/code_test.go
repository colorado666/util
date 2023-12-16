package example

import (
	"fmt"
	"gitee.com/asktop_golib/util/acoding"
	"testing"
)

func TestCoding(t *testing.T) {
	str := "月色真美，风也温柔，233333333，~！@#"                         //go字符串编码为utf-8
	fmt.Println("before convert:", str)                       //打印转换前的字符串
	fmt.Println("coding:", acoding.GetStrCoding([]byte(str))) //判断是否是utf-8

	gbkData := acoding.ConvertUTF8ToGBK([]byte(str))      //使用官方库将utf-8转换为gbk
	fmt.Println("gbk直接打印会出现乱码:", string(gbkData))         //乱码字符串
	fmt.Println("coding:", acoding.GetStrCoding(gbkData)) //判断是否是gbk

	utf8Data := acoding.ConvertGBKToUTF8(gbkData)          //将gbk再转换为utf-8
	fmt.Println("after convert:", string(utf8Data))        //打印转换后的字符串
	fmt.Println("coding:", acoding.GetStrCoding(utf8Data)) //判断是否是utf-8
}

//func TestCode(t *testing.T) {
//	str := "月色真美，风也温柔，233333333，~！@#"                         //go字符串编码为utf-8
//	fmt.Println("before convert:", str)                       //打印转换前的字符串
//	fmt.Println("coding:", acoding.GetStrCoding([]byte(str))) //判断是否是utf-8
//
//	gbkData := acoding.ConvertToString(str, acoding.CodeUTF8, acoding.CodeGBK)
//	fmt.Println(gbkData)
//	fmt.Println("coding:", acoding.GetStrCoding([]byte(gbkData))) //判断是否是gbk
//
//	utf8Data := acoding.ConvertToString(gbkData, acoding.CodeGBK, acoding.CodeUTF8)
//	fmt.Println(utf8Data)
//	fmt.Println("coding:", acoding.GetStrCoding([]byte(utf8Data))) //判断是否是utf-8
//}
