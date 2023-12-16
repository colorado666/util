package atime

import (
	"fmt"
	"gitee.com/asktop_golib/util/acast"
	"strings"
)

func toInt64(timestamp interface{}) int64 {
	return acast.ToInt64(timestamp)
}

// 秒级时间戳 转换为 毫秒级时间戳
func UnixToMilli(timestamp interface{}) int64 {
	return toInt64(timestamp) * 1e3
}

// 秒级时间戳 转换为 纳秒级时间戳
func UnixToNano(timestamp interface{}) int64 {
	return toInt64(timestamp) * 1e9
}

// 毫秒级时间戳 转换为 秒级时间戳
func MilliToUnix(timestamp interface{}) int64 {
	return toInt64(timestamp) / 1e3
}

// 毫秒级时间戳 转换为 纳秒时间戳
func MilliToNano(timestamp interface{}) int64 {
	return toInt64(timestamp) * 1e6
}

// 纳秒级时间戳 转换为 秒级时间戳
func NanoToUnix(timestamp interface{}) int64 {
	return toInt64(timestamp) / 1e9
}

// 纳秒级时间戳 转换为 毫秒级时间戳
func NanoToMilli(timestamp interface{}) int64 {
	return toInt64(timestamp) / 1e6
}

// 将秒级时间段转化成[XX天XX时XX分XX秒]
func UnixDurationToText(unixDuration int64, durationTexts string, force ...bool) string {
	var timeText string
	var isForce bool
	if len(force) > 0 {
		isForce = force[0]
	}
	if len(durationTexts) == 0 {
		durationTexts = "天,时,分,秒"
	}
	for _, durationText := range strings.Split(durationTexts, ",") {
		timeText, unixDuration = getUnixDurationToText(timeText, unixDuration, durationText, isForce)
	}
	return timeText
}

func getUnixDurationToText(timeText string, unixDuration int64, durationText string, force bool) (timeTextNew string, unixDurationNew int64) {
	var duration int64
	switch durationText {
	case "天":
		duration = 60 * 60 * 24
		number := unixDuration / duration
		if number > 0 {
			unixDuration -= number * duration
		}
		if force || number > 0 || timeText != "" {
			timeText = fmt.Sprintf("%s%d天", timeText, number)
		}
	case "时":
		duration = 60 * 60
		number := unixDuration / duration
		if number > 0 {
			unixDuration -= number * duration
		}
		if force || number > 0 || timeText != "" {
			timeText = fmt.Sprintf("%s%d时", timeText, number)
		}
	case "分":
		duration = 60
		number := unixDuration / duration
		if number > 0 {
			unixDuration -= number * duration
		}
		if force || number > 0 || timeText != "" {
			timeText = fmt.Sprintf("%s%d分", timeText, number)
		}
	default:
		duration = 1
		number := unixDuration / duration
		if number > 0 {
			unixDuration -= number * duration
		}
		timeText = fmt.Sprintf("%s%d秒", timeText, number)
	}
	return timeText, unixDuration
}
