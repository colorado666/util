package astring

import (
	"gitee.com/asktop_golib/util/acast"
	"gitee.com/asktop_golib/util/atime"
	"strings"
	"time"
)

//身份证号过滤处理
func FilterCardNo(cardNo string) string {
	return strings.ToUpper(strings.TrimSpace(cardNo))
}

//从身份证号取性别 1：男，2：女
func GetSexFromIDCard(cardNo string) int {
	if !IsIDCard(cardNo) {
		return 0
	}
	if len(cardNo) == 18 {
		sex := acast.ToInt(cardNo[16])
		if sex%2 == 0 {
			return 2
		} else {
			return 1
		}
	} else {
		sex := acast.ToInt(cardNo[13])
		if sex%2 == 0 {
			return 2
		} else {
			return 1
		}
	}
}

//从身份证号取生日
func GetBirthDayStrFromIDCard(cardNo string) string {
	cardNo = strings.TrimSpace(cardNo)
	if cardNo == "" {
		return ""
	}
	if !IsIDCard(cardNo) {
		return ""
	}
	var birthDayStr string
	if len(cardNo) == 18 {
		birthDayStr = cardNo[6:14]
	} else {
		birthDayStr = "19" + cardNo[6:12]
	}
	return birthDayStr[0:4] + "-" + birthDayStr[4:6] + "-" + birthDayStr[6:8]
}

//从身份证号取生日
func GetBirthDayInt64FromIDCard(cardNo string) int64 {
	birthDayStr := GetBirthDayStrFromIDCard(cardNo)
	if birthDayStr == "" {
		return 0
	}
	if birthDayStr == "1970-01-01" {
		return 1
	}
	birthDay, err := time.ParseInLocation("2006-01-02", birthDayStr, time.Local)
	if err != nil {
		return 0
	}
	return birthDay.Unix()
}

//从生日取年龄（周岁）
func GetAgeFromBirthDay(birthDay int64) int64 {
	if birthDay == 0 {
		return 0
	}
	birthDay = atime.StartDay(birthDay).Unix()
	age := (atime.StartDay().Unix() - birthDay) / (86400 * 365)
	return age
}

//从生日取年龄（虚岁）
func GetVirtualAgeFromBirthDay(birthDay int64) int64 {
	if birthDay == 0 {
		return 0
	}
	birthDay = atime.StartDay(birthDay).Unix()
	age := (atime.StartDay().Unix() - birthDay) / (86400 * 365)
	if (atime.StartDay().Unix()-birthDay)%(86400*365) > 0 {
		age++
	}
	return age
}

//从身份证号取年龄（周岁）
func GetAgeFromIDCard(cardNo string) int64 {
	birthDay := GetBirthDayInt64FromIDCard(cardNo)
	if birthDay == 0 {
		return 0
	}
	return GetAgeFromBirthDay(birthDay)
}

//从身份证号取年龄（虚岁）
func GetVirtualAgeFromIDCard(cardNo string) int64 {
	birthDay := GetBirthDayInt64FromIDCard(cardNo)
	if birthDay == 0 {
		return 0
	}
	return GetVirtualAgeFromBirthDay(birthDay)
}
