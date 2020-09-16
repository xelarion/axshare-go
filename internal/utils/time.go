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

func Strftime(dateTime time.Time, layout string) string {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return dateTime.In(loc).Format(layout)
}
