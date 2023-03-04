package timeUtil

import (
	"fmt"
	"time"
)

const (
	DayDuration    = 24 * time.Hour
	TwoDayDuration = 2 * DayDuration
)

func GetTimestampMillsecond() int64 {
	return time.Now().UnixMilli()
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

// 计算一个指定时间的前、后边界时间
// 这里的边界时间是指：以半天为单位，
// 如果是上午，beginTime为0点，engTime为12点
// 如果是上午，beginTime为12点，engTime为24点
// @param inTime 这个是millsecond, 13位整数
// return value 也是13位整数的 millsecond
func CalculateBoundaryTimeForHalfDay(inTime int64) (beginTime, endTime int64) {

	startHour := 0
	endHour := 24

	inputTime := time.UnixMilli(inTime)
	hour := inputTime.Hour()
	if (hour > 0) && (hour < 12) {
		startHour = 0
		endHour = 12
	} else {
		startHour = 12
		endHour = 24
	}

	year, month, day := inputTime.Date()

	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijingTimeZone := time.FixedZone("Beijing Time", secondsEastOfUTC)

	beginTmp := time.Date(year, month, day, startHour, 0, 0, 0, beijingTimeZone)
	endTmp := time.Date(year, month, day, endHour, 0, 0, 0, beijingTimeZone)

	beginTime = beginTmp.UnixMilli()
	endTime = endTmp.UnixMilli()

	return
}

// 计算一个指定时间的开始边界时间
// 这里的边界时间是指：以半天为单位，
// 如果是上午，beginTime为0点，engTime为12点
// 如果是上午，beginTime为12点，engTime为24点
// @param inTime 这个是millsecond, 13位整数
// return value 也是13位整数的 millsecond
func CalculateStartBoundaryTimeForHalfDay(inTime int64) (startBoundaryTime int64) {

	inputTime := time.UnixMilli(inTime)
	hour := inputTime.Hour()
	// if (hour > 0) && (hour < 12) {
	// 	startHour = 0
	// } else {
	// 	startHour = 12
	// }
	startHour := (hour / 12) * 12
	fmt.Println("startHour:", startHour)

	year, month, day := inputTime.Date()

	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijingTimeZone := time.FixedZone("Beijing Time", secondsEastOfUTC)

	beginTmp := time.Date(year, month, day, startHour, 0, 0, 0, beijingTimeZone)

	startBoundaryTime = beginTmp.UnixMilli()

	return
}
