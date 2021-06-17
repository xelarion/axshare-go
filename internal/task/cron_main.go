package task

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

// CronMain 定时任务
func CronMain() {
	c := cron.New()
	// 每天五点执行
	_, err := c.AddFunc("0 5 * * *", cleanAxureOldAttachments)
	if err != nil {
		logrus.Error("CronMain AddFunc error :", err)
	}

	c.Start()
}
