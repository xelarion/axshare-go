package utils

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func InitEnv() {
	// 从.env文件加载env变量
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func IsProductionEnv() bool {
	env := viper.GetString("env")
	return env == "production"
}
