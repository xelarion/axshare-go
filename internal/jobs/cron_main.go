package jobs

import (
	"github.com/robfig/cron"
)

// 定时任务
func CronMain() {
	c := cron.New()
	_ = c.AddFunc("@every 5d", CleanOldFiles)

	c.Start()
}
