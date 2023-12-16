package astring

import (
	"gitee.com/asktop_golib/util/acast"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

//字符长度
func Len(s string) int {
	return len([]rune(s))
}

//截取字符串
// @param length 不设置：截取全部；负数：向前截取
func Substr(s string, start int, length ...int) string {
	rs := []rune(s)
	l := len(rs)
	if len(length) > 0 {
		l = length[0]
	}
	if l > 0 {
		if start <= 0 {
			start = 0
		} else {
			if start > len(rs) {
				start = start % len(rs)
			}
		}

		end := start + l
		if start+l > len(rs) {
			end = len(rs)
		}
		return string(rs[start:end])
	} else if l < 0 {
		if start <= 0 {
			start = len(rs)
		} else {
			if start > len(rs) {
				start = start % len(rs)
			}
		}
		end := start

		start = end + l
		if end+l < 0 {
			start = 0
		}
		return string(rs[start:end])
	} else {
		return ""
	}
}

//截取字符串
// @param end 0：截取全部；负数：从后往前
func SubstrByEnd(s string, start int, end int) string {
	rs := []rune(s)

	if start < 0 {
		start = 0
	}
	if start > len(rs) {
		start = start % len(rs)
	}

	if end >= 0 {
		if end < start || end > len(rs) {
			end = len(rs)
		}
	} else {
		if len(rs)+end < start {
			end = len(rs)
		} else {
			end = len(rs) + end
		}
	}

	return string(rs[start:end])
}

//字符串是否相同（不区分大小写）
func EqualNoCase(str1 interface{}, str2 interface{}) bool {
	return strings.ToLower(acast.ToString(str1)) == strings.ToLower(acast.ToString(str2))
}

//替换字符串（不区分大小写）
func ReplaceNoCase(s string, old string, new string, n int) string {
	if n == 0 {
		return s
	}

	ls := strings.ToLower(s)
	lold := strings.ToLower(old)

	if m := strings.Count(ls, lold); m == 0 {
		return s
	} else if n < 0 || m < n {
		n = m
	}

	ns := make([]byte, len(s)+n*(len(new)-len(old)))
	w := 0
	start := 0
	for i := 0; i < n; i++ {
		j := start
		if len(old) == 0 {
			if i > 0 {
				_, wid := utf8.DecodeRuneInString(s[start:])
				j += wid
			}
		} else {
			j += strings.Index(ls[start:], lold)
		}
		w += copy(ns[w:], s[start:j])
		w += copy(ns[w:], new)
		start = j + len(old)
	}
	w += copy(ns[w:], s[start:])
	return string(ns[0:w])
}

//删除字符串两端的空格(含tab)，同时将中间多个空格(含tab)的转换为一个
func TrimSpaceToOne(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Replace(s, "	", " ", -1)      //替换tab为空格
	reg, _ := regexp.Compile("\\s{2,}")          //编译正则表达式
	s2 := make([]byte, len(s))                   //定义字符数组切片
	copy(s2, s)                                  //将字符串复制到切片
	spc_index := reg.FindStringIndex(string(s2)) //在字符串中搜索
	for len(spc_index) > 0 {                     //找到适配项
		s2 = append(s2[:spc_index[0]+1], s2[spc_index[1]:]...) //删除多余空格
		spc_index = reg.FindStringIndex(string(s2))            //继续在字符串中搜索
	}
	return string(s2)
}

// int 转换成指定长度的 string
// @param force 强制转换，当num长度大于length时，删除前面超过的部分
func IntToStr(num int, length int, force ...bool) string {
	if length <= 0 {
		return strconv.Itoa(num)
	} else {
		if num < 0 {
			numStr := strconv.Itoa(-num)
			if len(force) > 0 && force[0] || len(numStr) < length {
				numStr = strings.Repeat("0", length) + numStr
				return "-" + numStr[len(numStr)-length:]
			} else {
				return "-" + numStr
			}
		} else {
			numStr := strconv.Itoa(num)
			if len(force) > 0 && force[0] || len(numStr) < length {
				numStr = strings.Repeat("0", length) + numStr
				return numStr[len(numStr)-length:]
			} else {
				return numStr
			}
		}
	}
}

// int 转换成指定长度的 string
// @param force 强制转换，当num长度大于length时，删除前面超过的部分
func Int64ToStr(num int64, length int, force ...bool) string {
	if length <= 0 {
		return strconv.FormatInt(num, 10)
	} else {
		if num < 0 {
			numStr := strconv.FormatInt(-num, 10)
			if len(force) > 0 && force[0] || len(numStr) < length {
				numStr = strings.Repeat("0", length) + numStr
				return "-" + numStr[len(numStr)-length:]
			} else {
				return "-" + numStr
			}
		} else {
			numStr := strconv.FormatInt(num, 10)
			if len(force) > 0 && force[0] || len(numStr) < length {
				numStr = strings.Repeat("0", length) + numStr
				return numStr[len(numStr)-length:]
			} else {
				return numStr
			}
		}
	}
}

func Trim(str string, sep ...string) string {
	sepTemp := ","
	if len(sep) > 0 && len(sep[0]) > 0 {
		sepTemp = sep[0]
	}
	return strings.Trim(str, sepTemp)
}

func AddPrefix(str string, prefix string) string {
	str = strings.TrimPrefix(str, prefix)
	return prefix + str
}

func AddSuffix(str string, suffix string) string {
	str = strings.TrimSuffix(str, suffix)
	return str + suffix
}

func Add(str string, sep ...string) string {
	sepTemp := ","
	if len(sep) > 0 && len(sep[0]) > 0 {
		sepTemp = sep[0]
	}
	str = strings.Trim(str, sepTemp)
	if str == "" {
		return ""
	}
	return sepTemp + str + sepTemp
}

func Like(str string, sep ...string) string {
	if len(sep) > 0 && len(sep[0]) > 0 {
		str = Add(str, sep[0])
	}
	return Add(str, "%")
}

//将多个对象拼接成字符串
func Join(args ...interface{}) string {
	var rs string
	for _, arg := range args {
		rs += acast.ToStringForce(arg) + " "
	}
	return strings.TrimSpace(rs)
}

func JoinAdd(elems []string, sep ...string) string {
	sepTemp := ","
	if len(sep) > 0 && len(sep[0]) > 0 {
		sepTemp = sep[0]
	}
	str := strings.Join(elems, sepTemp)
	return Add(str, sepTemp)
}

func SplitTrim(str string, sep ...string) []string {
	sepTemp := ","
	if len(sep) > 0 && len(sep[0]) > 0 {
		sepTemp = sep[0]
	}
	str = strings.Trim(str, sepTemp)
	return strings.Split(str, sepTemp)
}
