package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"time"
)

var AxshareDb *gorm.DB

func InitDbConnection(dbKey string) {
	host := viper.GetString(dbKey + ".host")
	port := viper.GetString(dbKey + ".port")
	user := viper.GetString(dbKey + ".user")
	database := viper.GetString(dbKey + ".database")
	password := viper.GetString(dbKey + ".password")
	connectSql := user + ":" + password + "@tcp(" + host +
		":" + port + ")/" + database + "?" +
		"&charset=utf8mb4&parseTime=True&parseTime=True&loc=Local"
	axshareDb, err := gorm.Open("mysql", connectSql)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	// 设置更新数据库时间使用 UTC
	gorm.NowFunc = func() time.Time {
		return time.Now().UTC()
	}
	AxshareDb = axshareDb
}
