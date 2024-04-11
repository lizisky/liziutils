package timeUtil

import "time"

// const related with time
const (
	MillisecondPerMinute     = int64(60000)         // 一分钟的毫秒数
	MillisecondPerHour       = int64(3600000)       // 一小时的毫秒数
	MillisecondPerDay        = int64(86400000)      // 一天的毫秒数
	UnixMillisecond2023Begin = int64(1672531200000) // 以毫秒记录的时间戳，2023年1月1日0点
	UnixMillisecond2024Begin = int64(1704067200000) // 以毫秒记录的时间戳，2024年1月1日0点
)

const (
	DayDuration     = 24 * time.Hour
	TwoDaysDuration = 2 * DayDuration
)
