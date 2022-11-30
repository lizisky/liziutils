package utils

import "time"

const (
	DayDuration    = 24 * time.Hour
	TwoDayDuration = 2 * DayDuration
)

func GetTimestamp13() int64 {
	now := time.Now()
	return now.UnixNano() / 1000000
}

// 按 UTC 时间，判断 cur 是不是 last 的第二天
func IsNextDay(last, cur int64) bool {
	lastDate := GetDayStart(last)
	curDate := GetDayStart(cur)
	duration := curDate.Sub(lastDate)
	if (duration >= DayDuration) && (TwoDayDuration > duration) {
		return true
	} else {
		return false
	}
}

// 获取时间戳 UTC 时间的当日凌晨时间
func GetDayStart(timestamp int64) time.Time {
	timeUtc := Timestamp13ToDate(timestamp)
	year, month, day := timeUtc.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func Timestamp13ToDate(timestamp int64) time.Time {
	second := timestamp / 1000
	nanosecond := timestamp % 1000 * 1000000
	timeLocal := time.Unix(second, nanosecond)
	timeUtc := timeLocal.UTC()
	return timeUtc
}

func GetFormatDateNow() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05")
}
