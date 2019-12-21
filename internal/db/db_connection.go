package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"time"
)

var AxshareDb *gorm.DB

func InitDbConnection() {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	database := os.Getenv("MYSQL_DATABASE")
	password := os.Getenv("MYSQL_PASSWORD")
	connectSql := user + ":" + password + "@tcp(" + host +
		":" + port + ")/" + database + "?" +
		"&charset=utf8mb4&parseTime=True&parseTime=True&loc=Local"
	axshareDb, err := gorm.Open("mysql", connectSql)
	if err != nil {
		panic("failed to connect database." + err.Error())
	}
	// 设置更新数据库时间使用 UTC
	gorm.NowFunc = func() time.Time {
		return time.Now().UTC()
	}

	// 全局更新/删除锁
	axshareDb.BlockGlobalUpdate(true)
	axshareDb.DB().SetMaxIdleConns(0)
	axshareDb.DB().SetMaxOpenConns(500)
	AxshareDb = axshareDb
}
