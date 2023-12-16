package atime

import (
	"time"
)

//将 指定格式的时间字符串 转换成 纳秒级时间戳
func UnFormatUnixNanoTimestamp(format string, timeStr string) int64 {
	t, err := time.ParseInLocation(format, timeStr, time.Local)
	if err != nil {
		ErrLogFunc("gitee.com/asktop_golib/util/atime UnFormatUnixNanoTimestamp", "format:", format, "timeStr:", timeStr, "err:", err)
		return 0
	}
	return t.UnixNano()
}

//将 指定格式的时间字符串 转换成 毫秒级时间戳
func UnFormatMilliTimestamp(format string, timeStr string) int64 {
	t, err := time.ParseInLocation(format, timeStr, time.Local)
	if err != nil {
		ErrLogFunc("gitee.com/asktop_golib/util/atime UnFormatMilliTimestamp", "format:", format, "timeStr:", timeStr, "err:", err)
		return 0
	}
	return t.UnixNano() / 1e6
}

//将 指定格式的时间字符串 转换成 秒级时间戳
func UnFormatTimestamp(format string, timeStr string) int64 {
	t, err := time.ParseInLocation(format, timeStr, time.Local)
	if err != nil {
		ErrLogFunc("gitee.com/asktop_golib/util/atime UnFormatTimestamp", "format:", format, "timeStr:", timeStr, "err:", err)
		return 0
	}
	return t.Unix()
}

//将 指定格式的时间字符串 转换成 秒级时间戳
func UnFormatDateTime(timeStr string) int64 {
	return UnFormatTimestamp(DATETIME, timeStr)
}

//将 指定格式的时间字符串 转换成 秒级时间戳
func UnFormatDate(timeStr string) int64 {
	return UnFormatTimestamp(DATE, timeStr)
}

//将 指定格式的时间字符串 转换成 秒级时间戳
func UnFormatMonth(timeStr string) int64 {
	return UnFormatTimestamp(MONTH, timeStr)
}
