package akey

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

//hmac-md5单向秘钥key加密
func HmacMd5Byte(src []byte, key []byte) []byte {
	hash := hmac.New(md5.New, key)
	hash.Write(src)
	return hash.Sum(nil)
}

//hmac-md5单向秘钥key加密 32位
func HmacMd5(str, key string) string {
	hash := hmac.New(md5.New, []byte(key))
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func HmacMd5Base64(str, key string) string {
	return base64.StdEncoding.EncodeToString(HmacMd5Byte([]byte(str), []byte(key)))
}

//hmac-sha1单向秘钥key加密
func HmacSha1Byte(src []byte, key []byte) []byte {
	hash := hmac.New(sha1.New, key)
	hash.Write(src)
	return hash.Sum(nil)
}

//hmac-sha1单向秘钥key加密 40位
func HmacSha1(str, key string) string {
	hash := hmac.New(sha1.New, []byte(key))
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func HmacSha1Base64(str, key string) string {
	return base64.StdEncoding.EncodeToString(HmacSha1Byte([]byte(str), []byte(key)))
}

//hmac-sha256单向秘钥key加密
func HmacSha256Byte(src []byte, key []byte) []byte {
	hash := hmac.New(sha256.New, key)
	hash.Write(src)
	return hash.Sum(nil)
}

//hmac-sha256单向秘钥key加密 64位
func HmacSha256(str, key string) string {
	hash := hmac.New(sha256.New, []byte(key))
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func HmacSha256Base64(str, key string) string {
	return base64.StdEncoding.EncodeToString(HmacSha256Byte([]byte(str), []byte(key)))
}
