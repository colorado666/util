package akey

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
)

//md5单向加密
func Md5Byte(src []byte) []byte {
	hash := md5.New()
	hash.Write(src)
	return hash.Sum(nil)
}

//md5单向加密 32位
func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func Md5Hex(str string) string {
	return hex.EncodeToString(Md5Byte([]byte(str)))
}

func Md5Base64(str string) string {
	return base64.StdEncoding.EncodeToString(Md5Byte([]byte(str)))
}

//md5循环加密100次
func Md5Inum(str string, num int) (md5str string) {
	if num == 0 {
		num = 100
	}
	md5str = str
	for i := 1; i <= num; i++ {
		md5str = Md5(md5str)
	}
	return md5str
}

//随机md5
func RandMd5() string {
	data := make([]byte, 16)
	rand.Read(data)
	return hex.EncodeToString(data)
}
