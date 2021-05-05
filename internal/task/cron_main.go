package task

import (
	"github.com/robfig/cron"
)

// 定时任务
func CronMain() {
	c := cron.New()
	_ = c.AddFunc("@every 1d", cleanAxureOldAttachments)

	c.Start()
}
