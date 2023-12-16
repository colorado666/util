package acoding

import "golang.org/x/text/encoding/simplifiedchinese"

//将gbk转换为utf-8
func ConvertGBKToUTF8(gbkData []byte) []byte {
	utf8Data, _ := simplifiedchinese.GBK.NewDecoder().Bytes(gbkData)
	return utf8Data
}

//将utf-8转换为gbk
func ConvertUTF8ToGBK(utf8Data []byte) []byte {
	gbkData, _ := simplifiedchinese.GBK.NewEncoder().Bytes(utf8Data)
	return gbkData
}
