package main

import (
	"axshare_go/api"
	"axshare_go/internal/db"
	tasks "axshare_go/internal/task"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

const LogPath = "log/axshare_go.log"

func main() {
	initLogger()
	log.Info("axshare main start")

	initConfig()

	db.InitDbConnection("axshare_db")

	//db.AxshareDb.LogMode(true)

	serverChan := make(chan int)

	tasks.CronMain()
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
