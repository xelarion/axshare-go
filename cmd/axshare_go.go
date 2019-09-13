package main

import (
	"axshare_go/api"
	"axshare_go/internal/db"
	"axshare_go/internal/db/migrate"
	tasks "axshare_go/internal/task"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

const LogPath = "log/axshare_go.log"

func main() {
	initLogger()
	log.Info("axshare main start")

	initConfig()

	// 从.env文件加载env变量
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	db.InitDbConnection("axshare_db")
	migrate.Migrate()
	migrate.Seed()

	//db.AxshareDb.LogMode(true)

	serverChan := make(chan int)

	tasks.CronMain()

	if isProductionEnv() {
		gin.SetMode(gin.ReleaseMode)
	}

	go api.HttpServerRun()

	<-serverChan
}

func initConfig() {
	viper.AddConfigPath("configs")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func initLogger() {
	file, err := os.OpenFile(LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
	}
}

func isProductionEnv() bool {
	env := viper.GetString("env")
	return env == "production"
}
