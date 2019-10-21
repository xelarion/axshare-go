package migrate

import (
	"axshare_go/internal/db"
	"axshare_go/internal/models"
	"github.com/sirupsen/logrus"
)

func Migrate() {
	logrus.Info("migrate ...")
	db.AxshareDb.Set("gorm:table_options", "CHARSET=utf8mb4").Debug().AutoMigrate(
		&models.User{},
		&models.AxureGroup{},
		&models.Axure{},
		&models.Attachment{})
}

func Seed() {
	account := models.Account{Email: "admin@qq.com", Username: "admin", Password: "admin123456"}
	account.Create()
}
