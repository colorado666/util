package atime

import (
	"gitee.com/asktop_golib/util/acast"
	"time"
)

// 错误日志打印方法
var ErrLogFunc func(v ...interface{})

func init() {
	ErrLogFunc = func(v ...interface{}) {
		//默认不打印
	}
}

//将 纳秒级时间戳 转换成 本地时区时间
func ParseUnixNanoTimestamp(timestamp interface{}) (time.Time, error) {
	var err error
	fn := Now()
	sec, err := acast.ToInt64E(timestamp)
	if err != nil {
		return fn, err
	}
	nsec := sec % 1e9
	sec = sec / 1e9
	return time.Unix(sec, nsec), nil
}

//将 毫秒级时间戳 转换成 本地时区时间
func ParseMilliTimestamp(timestamp interface{}) (time.Time, error) {
	var err error
	fn := Now()
	sec, err := acast.ToInt64E(timestamp)
	if err != nil {
		return fn, err
	}
	nsec := (sec % 1e3) * 1e6
	sec = sec / 1e3
	return time.Unix(sec, nsec), nil
}

//将 秒级时间戳 转换成 本地时区时间
func ParseTimestamp(timestamp interface{}) (time.Time, error) {
	var err error
	fn := Now()
	sec, err := acast.ToInt64E(timestamp)
	if err != nil {
		return fn, err
	}
	return time.Unix(sec, 0), nil
}

//将 当前时间戳 转换成 指定格式的时间字符串
func FormatNow(format string) string {
	return Now().Format(format)
}

//将 纳秒级时间戳 转换成 指定格式的时间字符串
func FormatUnixNanoTimestamp(format string, timestamp interface{}) string {
	fn, err := ParseUnixNanoTimestamp(timestamp)
	if err != nil {
		if err.Error() != "" {
			ErrLogFunc("gitee.com/asktop_golib/util/atime ParseUnixNanoTimestamp", "timestamp:", timestamp, "err:", err)
		}
		return ""
	}
	return fn.Format(format)
}

//将 毫秒级时间戳 转换成 指定格式的时间字符串
func FormatMilliTimestamp(format string, timestamp interface{}) string {
	fn, err := ParseMilliTimestamp(timestamp)
	if err != nil {
		if err.Error() != "" {
			ErrLogFunc("gitee.com/asktop_golib/util/atime ParseMilliTimestamp", "timestamp:", timestamp, "err:", err)
		}
		return ""
	}
	return fn.Format(format)
}

//将 秒级时间戳 转换成 指定格式的时间字符串
func FormatTimestamp(format string, timestamp interface{}) string {
	fn, err := ParseTimestamp(timestamp)
	if err != nil {
		if err.Error() != "" {
			ErrLogFunc("gitee.com/asktop_golib/util/atime ParseTimestamp", "timestamp:", timestamp, "err:", err)
		}
		return ""
	}
	return fn.Format(format)
}

//将 秒级时间戳 转换成 指定格式的时间字符串 格式："2006-01-02 15:04:05"
func FormatDateTime(timestamp interface{}) string {
	return FormatTimestamp(DATETIME, timestamp)
}

//将 秒级时间戳 转换成 指定格式的时间字符串 格式："2006-01-02"
func FormatDate(timestamp interface{}) string {
	return FormatTimestamp(DATE, timestamp)
}

//将 秒级时间戳 转换成 指定格式的时间字符串 格式："15:04:05"
func FormatTime(timestamp interface{}) string {
	return FormatTimestamp(TIME, timestamp)
}

//将 秒级时间戳 转换成 指定格式的时间字符串 格式："2006-01"
func FormatMonth(timestamp interface{}) string {
	return FormatTimestamp(MONTH, timestamp)
}
