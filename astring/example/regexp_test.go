package example

import (
	"gitee.com/asktop_golib/util/astring"
	"regexp"
	"strconv"
	"testing"
)

func TestIsNum_EN(t *testing.T) {
	t.Log(astring.IsNum_Alpha("jasdf_2934_"))
	t.Log(astring.IsNum_Alpha("jasdf_29.34_"))
	t.Log(astring.HasNum_Alpha("jasdf-29.34"))
}

func TestIsUpperChar(t *testing.T) {
	a := "a"
	b := "A"
	c := "Aa"
	t.Log(astring.IsUpper(a))
	t.Log(astring.IsUpper(b))
	t.Log(astring.IsUpper(c))

	t.Log("----------")

	t.Log(astring.HasUpper(a))
	t.Log(astring.HasUpper(b))
	t.Log(astring.HasUpper(c))

	t.Log("----------")

	t.Log(astring.IsLower(a))
	t.Log(astring.IsLower(b))
	t.Log(astring.IsLower(c))

	t.Log("----------")

	t.Log(astring.HasLower(a))
	t.Log(astring.HasLower(b))
	t.Log(astring.HasLower(c))
}

func TestIsAllDecimal(t *testing.T) {
	t.Log(astring.IsDecimal("213123_2934"))
	t.Log(astring.IsDecimal("-213123.2934"))
	t.Log(astring.IsDecimal("213123.293400", 8))
	t.Log(astring.IsDecimal("-213123.293400", 6))
	t.Log(astring.IsDecimal("213123.293400", 3, 8))
	t.Log(astring.IsDecimal("-213123.293400", 3, 6))
	t.Log(astring.IsDecimal("-213123", 0))
	t.Log(astring.IsDecimal("-213123", 0, 1))
	t.Log(astring.IsDecimal("-213123.", 0, 1))
	t.Log(astring.IsDecimal("213123"))
}

func TestIsDateFormat(t *testing.T) {
	data1 := "2018-2-1"
	data2 := "1918.01.1"
	data3 := "2018年12月30日"
	data4 := "2018/10/1"
	t.Log(astring.IsDateFormat(data1, "-"))
	t.Log(astring.IsDateFormat(data2, "."))
	t.Log(astring.IsDateFormat(data3, "/"))
	t.Log(astring.IsDateFormat(data4, "/"))
}

//匹配并替换
func TestRegReplace(t *testing.T) {
	data := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pattern := "[0-9]+.[0-9]+"
	repl := "##.#"

	//将匹配到的浮点数替换为"##.#"
	str := regexp.MustCompile(pattern).ReplaceAllString(data, repl)
	t.Log(str)

	//将匹配到的浮点数替换为乘以2的浮点数
	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}
	str2 := regexp.MustCompile(pattern).ReplaceAllStringFunc(data, f)
	t.Log(str2)
}

//匹配并获取第一条匹配结果
func TestFindString(t *testing.T) {
	pattern := "共(\\d+)页"
	data := `<td height="2">共12页&nbsp;1300条&nbsp;共23页<u>首页</u>&nbsp;<u>上一页</u>&nbsp;&nbsp;<a href=ClassList-42-1.html><u><font color=red>1</u></font></a>&nbsp;&nbsp;<a href=ClassList-42-2.html><u>2</u></a>&nbsp;&nbsp;<a href=ClassList-42-3.html><u>3</u></a>&nbsp;&nbsp;<a href=ClassList-42-4.html><u>4</u></a>&nbsp;&nbsp;<a href=ClassList-42-5.html><u>5</u></a>&nbsp;&nbsp;<a href=ClassList-42-6.html><u>6</u></a>&nbsp;&nbsp;<a href=ClassList-42-7.html><u>7</u></a>&nbsp;&nbsp;<a href=ClassList-42-8.html><u>8</u></a>&nbsp;&nbsp;<a href=ClassList-42-9.html><u>9</u></a>&nbsp;&nbsp;<a href=ClassList-42-10.html><u>10</u></a>&nbsp;&nbsp;<a href=ClassList-42-2.html><u>下一页</u></a>&nbsp;<a href=ClassList-42-12.html><u>尾页</u></a>&nbsp;114条/页&nbsp;</td>`
	result1 := regexp.MustCompile(pattern).FindString(data)
	t.Log(result1)

	//查看与上面的区别，若想提取出正则匹配结果，必须加括号
	result2 := regexp.MustCompile(pattern).FindStringSubmatch(data)
	t.Log(result2)
	t.Log(result2[1])
}

//匹配并获取所有匹配结果
func TestFindAllString(t *testing.T) {
	pattern := "共(\\d+)页"
	data := `<td height="2">共12页&nbsp;1300条&nbsp;共23页<u>首页</u>&nbsp;<u>上一页</u>&nbsp;&nbsp;<a href=ClassList-42-1.html><u><font color=red>1</u></font></a>&nbsp;&nbsp;<a href=ClassList-42-2.html><u>2</u></a>&nbsp;&nbsp;<a href=ClassList-42-3.html><u>3</u></a>&nbsp;&nbsp;<a href=ClassList-42-4.html><u>4</u></a>&nbsp;&nbsp;<a href=ClassList-42-5.html><u>5</u></a>&nbsp;&nbsp;<a href=ClassList-42-6.html><u>6</u></a>&nbsp;&nbsp;<a href=ClassList-42-7.html><u>7</u></a>&nbsp;&nbsp;<a href=ClassList-42-8.html><u>8</u></a>&nbsp;&nbsp;<a href=ClassList-42-9.html><u>9</u></a>&nbsp;&nbsp;<a href=ClassList-42-10.html><u>10</u></a>&nbsp;&nbsp;<a href=ClassList-42-2.html><u>下一页</u></a>&nbsp;<a href=ClassList-42-12.html><u>尾页</u></a>&nbsp;114条/页&nbsp;</td>`
	//匹配并获取所有匹配结果，n为获取匹配的个数，n为-1时获取所有匹配结果
	result1 := regexp.MustCompile(pattern).FindAllString(data, 1)
	for i, result := range result1 {
		t.Log(i, ":", result)
	}

	//查看与上面的区别，若想提取出正则匹配结果，必须加括号
	result2 := regexp.MustCompile(pattern).FindAllStringSubmatch(data, -1)
	for i, result := range result2 {
		t.Log(i, ":", result)
		t.Log(result[1])
	}
}
