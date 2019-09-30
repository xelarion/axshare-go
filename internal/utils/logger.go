package utils

import (
	"github.com/sirupsen/logrus"
	"os"
)

const axshareLogPath = "log/axshare_go.log"

func InitLogger() {
	file, err := os.OpenFile(axshareLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		logrus.SetOutput(file)
	}
}
