package main

import (
	"axshare_go/api"
	"axshare_go/internal/db"
	"axshare_go/internal/db/migrate"
	"axshare_go/internal/jobs"
	"axshare_go/internal/models"
	"axshare_go/internal/task"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	initLogger()
	InitEnv()
	initDB()

	models.InitCacheConfig()

	serverChan := make(chan int)

	go jobs.CronMain()
	go api.RunHttpServer()
	go task.RunMachineryServer()

	<-serverChan
}

func initLogger() {
	file, err := os.OpenFile("log/production.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		logrus.SetOutput(file)
	}
}

func initDB() {
	db.InitDbConnection()
	migrate.Migrate()
	migrate.Seed()
}

func InitEnv() {
	// 从.env文件加载env变量
	err := godotenv.Load("config/.env")
	if err != nil {
		panic(err)
	}
}
