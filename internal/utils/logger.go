package utils

import (
	"github.com/sirupsen/logrus"
	"os"
)

func InitLogger() {
	file, err := os.OpenFile("log/production.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		logrus.SetOutput(file)
	}
}
