package example

import (
	"gitee.com/asktop_golib/util/astring"
	"testing"
)

func TestHideNo(t *testing.T) {
	t.Log(astring.HideNo("", 1, 1))
	t.Log(astring.HideNo("asw", 1, 1))
	t.Log(astring.HideNo("asw", 0, 1))
	t.Log(astring.HideNo("asw", 2, 2))
	t.Log(astring.HideNo("", 1, 1, 6))
	t.Log(astring.HideNo("a", 2, 2, 6))
	t.Log(astring.HideNo("as", 2, 2, 6))
	t.Log(astring.HideNo("asw", 1, 1, 6))
	t.Log(astring.HideNo("asdfjkhksdfkj", 3, 0, 6))
	t.Log(astring.HideNo("123456789", 2, 2, 6))
}

func TestHidePhone(t *testing.T) {
	t.Log(astring.HidePhone(""))
	t.Log(astring.HidePhone("13412345678"))
	t.Log(astring.HidePhone("+8613412345678"))
	t.Log(astring.HidePhone("8941123"))
	t.Log(astring.HidePhone("89414567"))
	t.Log(astring.HidePhone("0539-89414567"))
}

func TestHideEmail(t *testing.T) {
	t.Log(astring.HideEmail(""))
	t.Log(astring.HideEmail("as@163.com"))
	t.Log(astring.HideEmail("asdfjkhksdfkj@163.com"))
	t.Log(astring.HideEmail("123456789@163.com"))
	t.Log(astring.HideEmail("13412345678@163.com"))
}

func TestHidePwd(t *testing.T) {
	t.Log(astring.HidePwd(""))
	t.Log(astring.HidePwd("", true))
	t.Log(astring.HidePwd("asw"))
	t.Log(astring.HidePwd("123456789"))
}

func TestHideName(t *testing.T) {
	t.Log(astring.HideName(""))
	t.Log(astring.HideName("张"))
	t.Log(astring.HideName("张三"))
	t.Log(astring.HideName("张小三"))
	t.Log(astring.HideName("张小三四五六"))
}

func TestHideIdNo(t *testing.T) {
	t.Log(astring.HideIdNo(""))
	t.Log(astring.HideIdNo("11010119900307133X"))
	t.Log(astring.HideIdNo("110101199003072252"))
	t.Log(astring.HideIdNo("110101900307133"))
	t.Log(astring.HideIdNo("110101900307225"))
}

func TestHideBankNo(t *testing.T) {
	t.Log(astring.HideBankNo(""))
	t.Log(astring.HideBankNo("6222123456789098765"))
}

func TestHideAddress(t *testing.T) {
	t.Log(astring.HideAddress(""))
	t.Log(astring.HideAddress("山东省"))
	t.Log(astring.HideAddress("山东省临沂市"))
	t.Log(astring.HideAddress("山东省临沂市兰山区"))
	t.Log(astring.HideAddress("山东省临沂市兰山区北城新区"))
	t.Log(astring.HideAddress("山东省临沂市兰山区北城新区红日大厦"))
	t.Log(astring.HideAddress("山东省临沂市兰山区北城新区红日大厦15楼"))
	t.Log(astring.HideAddress("山东省临沂市兰山区北城新区红日大厦15楼1单元1502"))
}
