package astring

import (
	"strings"
)

// 隐藏字符串
// start：前端显示长度
// end：后端显示长度
// length：指定显示总长度，若不指定，则按原字符串长度输出
func HideNo(s string, start int, end int, length ...int) string {
	s = strings.TrimSpace(s)
	oldLen := len([]rune(s))
	if oldLen == 0 {
		return ""
	}

	newLen := oldLen
	if len(length) > 0 && length[0] > 0 {
		newLen = length[0]
	}

	minLen := oldLen
	if newLen < minLen {
		minLen = newLen
	}

	if minLen <= 1 {
		return strings.Repeat("*", newLen)
	}

	subStart := true
	for {
		if start+end >= minLen {
			if subStart {
				start--
				subStart = false
				if start < 0 {
					start = 0
				}
			} else {
				end--
				subStart = true
				if end < 0 {
					end = 0
				}
			}
		} else {
			break
		}
	}
	rs := Substr(s, 0, start) + strings.Repeat("*", newLen-start-end) + Substr(s, 0, -end)
	return rs
}

// 隐藏 手机号
func HidePhone(s string) string {
	s = strings.TrimSpace(s)
	length := len(s)
	if length == 0 {
		return ""
	}
	if strings.Contains(s, "+") {
		return Substr(s, 0, length-8) + "****" + SubstrByEnd(s, length-4, 0)
	} else {
		if strings.Contains(s, "-") || strings.Contains(s, "_") || strings.Contains(s, " ") {
			return Substr(s, 0, length-6) + "***" + SubstrByEnd(s, length-3, 0)
		} else {
			if length == 11 {
				return Substr(s, 0, 3) + "****" + SubstrByEnd(s, length-4, 0)
			} else {
				return Substr(s, 0, length-6) + "***" + SubstrByEnd(s, length-3, 0)
			}
		}
	}
}

// 隐藏 邮箱
func HideEmail(s string) string {
	emails := strings.Split(s, "@")
	if len(emails) != 2 {
		return s
	}
	return HideNo(emails[0], 2, 2, 6) + "@" + emails[1]
}

// 隐藏 密码
func HidePwd(s string, allHide ...bool) string {
	s = strings.TrimSpace(s)
	if len(allHide) > 0 && allHide[0] {
		return "******"
	} else {
		if len(s) > 0 {
			return "******"
		} else {
			return ""
		}
	}
}

// 隐藏 姓名
// 两个字的显示后1位，多余两个字的显示前1位后1位
func HideName(s string) string {
	if len(s) > 2 {
		return HideNo(s, 1, 1)
	} else {
		return HideNo(s, 0, 1)
	}
}

// 隐藏 身份证号
// 显示前6位，后4位
func HideIdNo(s string) string {
	return HideNo(s, 6, 4)
}

// 隐藏 银行卡号
// 显示前4位，后4位
func HideBankNo(s string) string {
	return HideNo(s, 4, 4)
}

// 隐藏 地址
func HideAddress(s string) string {
	strRune := []rune(s)
	n := len(strRune) / 3
	if n > 4 {
		n = 4
	}
	return Substr(s, 0, n) + strings.Repeat("*", len(strRune)-n*2) + SubstrByEnd(s, len(strRune)-n, 0)
}
