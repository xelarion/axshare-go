package migrate

import (
	"axshare_go/internal/db"
	"axshare_go/internal/models"
	"github.com/sirupsen/logrus"
	"github.com/xandercheung/acct"
)

func Migrate() {
	_ = db.AxshareDb.Set("gorm:table_options", "CHARSET=utf8mb4").Debug().AutoMigrate(
		&acct.Account{},
		&models.Config{},
		&models.AxureGroup{},
		&models.Axure{},
		&models.Attachment{})
}

func Seed() {
	if !acct.Finder.IsAccountWithDeletedExists(map[string]interface{}{"username": "admin"}, nil) {
		account := acct.Account{
			Email:    "admin@qq.com",
			Username: "admin",
			Nickname: "Admin",
			Password: "admin@123456"}
		if err := account.Create(); err != nil {
			logrus.Error("db seed error: ", err.Error())
		}
	}
}
