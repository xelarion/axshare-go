package main

import (
	"axshare_go/api"
	"axshare_go/internal/db"
	"axshare_go/internal/db/migrate"
	"axshare_go/internal/jobs"
	"axshare_go/internal/task"
	"axshare_go/internal/utils"
)

func main() {
	initLogger()
	initConfigEnv()
	initDB()

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
	//utils.InitConfig()
}

func initDB() {
	db.InitDbConnection()
	migrate.Migrate()
	migrate.Seed()
}
