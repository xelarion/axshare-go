package db

import (
	"github.com/xandercheung/acct"
	"gorm.io/gorm"
)

var AxshareDb *gorm.DB

func InitDbConnection() {
	dsn := acct.GetMysqlConnectArgsFromEnv()
	if err := acct.ConnectDB(dsn); err != nil {
		panic("failed to connect database." + err.Error())
	}

	AxshareDb = acct.DB
}
