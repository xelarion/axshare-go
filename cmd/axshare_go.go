package main

import (
	"axshare_go/api"
	"axshare_go/internal/db"
	"axshare_go/internal/db/migrate"
	"axshare_go/internal/jobs"
	"axshare_go/internal/task"
	"axshare_go/internal/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	initLogger()
	initConfigEnv()
	initDB()
	initGinSetting()

	serverChan := make(chan int)

	jobs.CronMain()
	go api.RunHttpServer()
	go task.RunMachineryServer()

	<-serverChan
}

func initLogger() {
	utils.InitLogger()
}

func initConfigEnv() {
	utils.InitEnv()
	utils.InitConfig()
}

func initDB() {
	db.InitDbConnection()
	migrate.Migrate()
	migrate.Seed()
}

func initGinSetting() {
	if utils.IsProductionEnv() {
		gin.SetMode(gin.ReleaseMode)
	}
}
