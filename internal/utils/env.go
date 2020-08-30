package utils

import (
	"github.com/joho/godotenv"
)

func InitEnv() {
	// 从.env文件加载env变量
	err := godotenv.Load("config/machinery.env", "config/axshare.env", "config/mysql.env")
	if err != nil {
		panic(err)
	}
}

