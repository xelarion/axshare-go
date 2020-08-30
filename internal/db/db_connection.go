package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/xandercheung/acct"
)

var AxshareDb *gorm.DB

func InitDbConnection() {
	var err error
	AxshareDb, err = acct.InitDBConnection("mysql", acct.GetMysqlConnectArgsFromEnv())
	if err != nil {
		panic("failed to connect database." + err.Error())
	}
}
