// Time 相关的函数
package utils

import (
	"time"
)

func FormatDateTime(dateTime time.Time) string {
	if !dateTime.IsZero() {
		loc, _ := time.LoadLocation("Asia/Shanghai")
		timeFormat := "2006-01-02 15:04:05"
		return dateTime.In(loc).Format(timeFormat)
	} else {
		return ""
	}
}

func FormatDateTimeNoMinute(dateTime time.Time) string {
	if !dateTime.IsZero() {
		loc, _ := time.LoadLocation("Asia/Shanghai")
		timeFormat := "2006-01-02"
		return dateTime.In(loc).Format(timeFormat)
	} else {
		return ""
	}
}

func FormatUtcDateTime(dateTime time.Time) string {
	if !dateTime.IsZero() {
		loc, _ := time.LoadLocation("UTC")
		timeFormat := "2006-01-02 15:04:05.999999"
		return dateTime.In(loc).Format(timeFormat)
	} else {
		return ""
	}
}

// 格式化返回 YYYYMM。例: 201906
func FormatDateTimeMonth(dateTime time.Time) string {
	if !dateTime.IsZero() {
		loc, _ := time.LoadLocation("Asia/Shanghai")
		timeFormat := "200601"
		return dateTime.In(loc).Format(timeFormat)
	} else {
		return ""
	}
}

func Strftime(dateTime time.Time, layout string) string {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return dateTime.In(loc).Format(layout)
}
