package example

import (
	"errors"
	"gitee.com/asktop_golib/util/astring"
	"testing"
)

func TestSubstr(t *testing.T) {
	str := "0123456789"
	t.Log(astring.Substr(str, 8))
	t.Log(astring.Substr(str, 8, 3))
	t.Log(astring.Substr(str, 0, 3))
	t.Log(astring.Substr(str, 0, -3))
	t.Log(astring.Substr(str, 1, -3))
}

func TestTrimSpaceToOne(t *testing.T) {
	t.Log("---" + astring.TrimSpaceToOne("\ta	b	  c  d	") + "---")
	t.Log("---" + astring.TrimSpaceToOne("a\t b	  c  d ") + "---")
}

func TestIntToStr(t *testing.T) {
	t.Log(astring.IntToStr(123, 6))
	t.Log(astring.IntToStr(123456789, 6))
	t.Log(astring.IntToStr(123456789, 6, true))
}

func TestJoin(t *testing.T) {
	a := map[string]interface{}{}
	a["a"] = "abc"
	a["b"] = 123
	e := errors.New("err")
	t.Log(astring.Join("uid:", 111, "data:", a, e))
}

func TestToFirstUpper(t *testing.T) {
	a := "abcdef"
	t.Log(astring.ToFirstUpper(a))
	b := "ABCDEF"
	t.Log(astring.ToFirstLower(b))
}

func TestToCamelCase(t *testing.T) {
	a := "user_id"
	b := "User_Id"
	c := "userId"
	d := "UserId"
	t.Log(astring.ToCamelUpper(a))
	t.Log(astring.ToCamelUpper(b))
	t.Log(astring.ToCamelUpper(c))
	t.Log(astring.ToCamelUpper(d))

	t.Log("----------")

	t.Log(astring.ToCamelLower(a))
	t.Log(astring.ToCamelLower(b))
	t.Log(astring.ToCamelLower(c))
	t.Log(astring.ToCamelLower(d))

	t.Log("----------")

	t.Log(astring.ToSnakeUpper(a))
	t.Log(astring.ToSnakeUpper(b))
	t.Log(astring.ToSnakeUpper(c))
	t.Log(astring.ToSnakeUpper(d))

	t.Log("----------")

	t.Log(astring.ToSnakeLower(a))
	t.Log(astring.ToSnakeLower(b))
	t.Log(astring.ToSnakeLower(c))
	t.Log(astring.ToSnakeLower(d))
}
