package acoding

//import (
//	"github.com/axgle/mahonia"
//)
//
//const (
//	CodeGBK  = "gbk"
//	CodeUTF8 = "utf-8"
//	CodeISO  = "iso-8859-1"
//)
//
//// 转码
//func ConvertToString(src string, srcCode string, tagCode string) string {
//	srcCoder := mahonia.NewDecoder(srcCode)
//	srcResult := srcCoder.ConvertString(src)
//	tagCoder := mahonia.NewDecoder(tagCode)
//	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
//
//	result := string(cdata)
//	return result
//}
