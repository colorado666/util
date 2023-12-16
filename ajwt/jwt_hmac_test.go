package ajwt

import (
	"fmt"
	"gitee.com/asktop_golib/util/atime"
	"testing"
	"time"
)

var (
	// 哈希时用的秘钥字符串
	keyStr = `1234567890`
)

// 生成jwt
func TestCreateJwtHmac(t *testing.T) {
	// exp := atime.Now().Unix()
	exp := atime.Now().Add(time.Second * 60 * 60 * 24).Unix()

	params := map[string]interface{}{
		"user_id":   "06d0bdf1-8393-49a3-4484-f585e1c42fb7",
		"user_name": "888888",
		"device":    "android",
	}

	jwtStr, err := CreateJwtHmac(params, exp, keyStr)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		// jwtStr = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzUyMzI1NjMsInVzZXJfaWQiOjEyM30.YsxomqFnmLWfa57PtWR-B32GZhx5dQdH1aLOb_CnjMQ`
		fmt.Println(jwtStr)
	}

	fmt.Println("--------------------")

	// 解析jwt
	info, err := ParseJwtHmac(jwtStr, keyStr)
	fmt.Println("IsExpired：", IsExpired(err))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(info)
	}

	fmt.Println("--------------------")

	// 解析jwt到结构
	obj := struct {
		UserId   string `json:"user_id"`
		UserName string `json:"user_name"`
		Device   string `json:"device"`
	}{}
	err = DecryptObj(jwtStr, &obj, keyStr)
	fmt.Println("IsExpired：", IsExpired(err))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(obj)
	}
}
