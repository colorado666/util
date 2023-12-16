package atime

import (
	"sync"
	"time"
)

const (
	UTCDATETIME = "2006-01-02T15:04:05Z"
	DATETIME    = "2006-01-02 15:04:05"
	DATE        = "2006-01-02"
	TIME        = "15:04:05"
	MONTH       = "2006-01"
)

var mu sync.RWMutex
var offsetTime time.Duration //时间偏移量
var fixedTime *time.Time     //固定时间

//设置时间偏量（改变当前时间）
func Offset(offset time.Duration) {
	mu.Lock()
	defer mu.Unlock()
	offsetTime = offset
}

//设置固定时间（改变当前时间）
func Fixed(fixed *time.Time) {
	mu.Lock()
	defer mu.Unlock()
	fixedTime = fixed
}

//当前时间
func Now() time.Time {
	mu.Lock()
	defer mu.Unlock()
	if fixedTime != nil {
		return *fixedTime
	}
	return time.Now().Add(offsetTime)
}

//获取当前时间 秒级时间戳
func NowUnix() int64 {
	return Now().Unix()
}

//获取当前时间 毫秒级时间戳
func NowMilli() int64 {
	return Now().UnixNano() / 1e6
}

//获取当前时间 纳秒级时间戳
func NowNano() int64 {
	return Now().UnixNano()
}
