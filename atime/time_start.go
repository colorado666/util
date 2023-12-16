package atime

import (
	"time"
)

const (
	Day  = time.Duration(time.Hour * 24)
	Week = Day * 7
)

func startTime(timestamp ...interface{}) time.Time {
	fn := Now()
	if len(timestamp) > 0 {
		var err error
		fn, err = ParseTimestamp(timestamp[0])
		if err != nil {
			ErrLogFunc("gitee.com/asktop_golib/util/atime ParseTimestamp", "timestamp:", timestamp, "err:", err)
		}
	}
	return fn
}

//获取 当前时间 或 指定时间戳 的 当年开始时间
func StartYear(timestamp ...interface{}) time.Time {
	fn := startTime(timestamp...)
	return time.Date(fn.Year(), 1, 1, 0, 0, 0, 0, time.Local)
}

//获取 当前时间 或 指定时间戳 的 上年开始时间
func StartYearLast(timestamp ...interface{}) time.Time {
	fn := startTime(timestamp...)
	day := time.Date(fn.Year(), 1, 1, 0, 0, 0, 0, time.Local)
	return day.AddDate(-1, 0, 0)
}

//获取 当前时间 或 指定时间戳 的 下年开始时间
func StartYearNext(timestamp ...interface{}) time.Time {
	fn := startTime(timestamp...)
	day := time.Date(fn.Year(), 1, 1, 0, 0, 0, 0, time.Local)
	return day.AddDate(1, 0, 0)
}

//获取 当前时间 或 指定时间戳 的 当前月开始时间
func StartMonth(timestamp ...interface{}) time.Time {
	fn := startTime(timestamp...)
	return time.Date(fn.Year(), fn.Month(), 1, 0, 0, 0, 0, time.Local)
}

//获取 当前时间 或 指定时间戳 的 上月开始时间
func StartMonthLast(timestamp ...interface{}) time.Time {
	fn := startTime(timestamp...)
	day := time.Date(fn.Year(), fn.Month(), 1, 0, 0, 0, 0, time.Local)
	return day.AddDate(0, -1, 0)
}

//获取 当前时间 或 指定时间戳 的 下月开始时间
func StartMonthNext(timestamp ...interface{}) time.Time {
	fn := startTime(timestamp...)
	day := time.Date(fn.Year(), fn.Month(), 1, 0, 0, 0, 0, time.Local)
	return day.AddDate(0, 1, 0)
}

//获取 当前时间 或 指定时间戳 的 当前周一时间
func StartWeek(timestamp ...interface{}) time.Time {
	fn := StartDay(timestamp...)
	offset := int(time.Monday - fn.Weekday())
	if offset > 0 {
		offset = -6
	}
	return fn.AddDate(0, 0, offset)
}

//获取 当前时间 或 指定时间戳 的 当天开始时间
func StartDay(timestamp ...interface{}) time.Time {
	fn := startTime(timestamp...)
	return time.Date(fn.Year(), fn.Month(), fn.Day(), 0, 0, 0, 0, time.Local)
}

//获取 当前时间 或 指定时间戳 的 当前小时开始时间
func StartHour(timestamp ...interface{}) time.Time {
	fn := startTime(timestamp...)
	return time.Date(fn.Year(), fn.Month(), fn.Day(), fn.Hour(), 0, 0, 0, time.Local)
}

//获取 当前时间 或 指定时间戳 的 当前分钟开始时间
func StartMinute(timestamp ...interface{}) time.Time {
	fn := startTime(timestamp...)
	return time.Date(fn.Year(), fn.Month(), fn.Day(), fn.Hour(), fn.Minute(), 0, 0, time.Local)
}
