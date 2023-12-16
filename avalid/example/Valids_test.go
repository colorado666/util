package example

import (
	"fmt"
	"gitee.com/asktop_golib/util/avalid"
	"testing"
)

func TestValid(t *testing.T) {
	if msg, ok := avalid.New("手机号", "").IsMobile().Check(); !ok {
		fmt.Println(msg)
	} else {
		fmt.Println(ok)
	}
	if msg, ok := avalid.New("手机号", "").Required().IsMobile().Check(); !ok {
		fmt.Println(msg)
	} else {
		fmt.Println(ok)
	}
	if msg, ok := avalid.New("手机号", "").IsMobile().Required().Check(); !ok {
		fmt.Println(msg)
	} else {
		fmt.Println(ok)
	}
}

func TestNew(t *testing.T) {
	if msg, ok := avalid.New("用户名", "abcd").Required("自定义返回消息：用户名必须").Check(); !ok {
		fmt.Println(msg)
	} else {
		fmt.Println(ok)
	}
	if msg, ok := avalid.New("用户名", "abcd").Required().Length(6, 20).Check(); !ok {
		fmt.Println(msg)
	} else {
		fmt.Println(ok)
	}
	if msg, ok := avalid.New("用户名", "abcd").Required().Same("abc").Check(); !ok {
		fmt.Println(msg)
	} else {
		fmt.Println(ok)
	}
	if msg, ok := avalid.New("用户名", "abcd").Required().InSlice([]string{"a", "b", "c", "abc"}).Check(); !ok {
		fmt.Println(msg)
	} else {
		fmt.Println(ok)
	}
	if msg, ok := avalid.New("金额", "12.3456").Required().IsInt("自定义返回消息：金额必须是整数").Check(); !ok {
		fmt.Println(msg)
	} else {
		fmt.Println(ok)
	}
	if msg, ok := avalid.New("金额", "12.3456").Required().IsDecimal(nil).Check(); !ok {
		fmt.Println(msg)
	} else {
		fmt.Println(ok)
	}
	if msg, ok := avalid.New("金额", "12.3456").Required().Between(12, "12.15").Check(); !ok {
		fmt.Println(msg)
	} else {
		fmt.Println(ok)
	}
}

func TestNews(t *testing.T) {
	valid := avalid.News().
		Valid("username", "abcd", "用户名").Required().Length(6, 20).
		Valid("amount", "12.3456").Required().Between(12, "12.15")

	if msg, ok := valid.Check(); !ok {
		fmt.Println(msg)
	} else {
		fmt.Println(ok)
	}

	msgs, ok := valid.Checks()
	if !ok {
		for k, v := range msgs {
			fmt.Println(k, v)
		}
	}
}
