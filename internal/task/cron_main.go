package task

import (
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
)

// 定时任务
func CronMain() {
	c := cron.New()
	// 每天五点执行
	err := c.AddFunc("0 5 * * *", cleanAxureOldAttachments)
	if err != nil {
		logrus.Error("CronMain AddFunc error :", err)
	}

	c.Start()
}
