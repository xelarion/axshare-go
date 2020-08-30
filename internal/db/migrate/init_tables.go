package migrate

import (
	"axshare_go/internal/db"
	"axshare_go/internal/models"
	"github.com/sirupsen/logrus"
	"github.com/xandercheung/acct"
)

func Migrate() {
	logrus.Info("migrate ...")
	acct.MigrateTables()
	db.AxshareDb.Set("gorm:table_options", "CHARSET=utf8mb4").Debug().AutoMigrate(
		&models.User{},
		&models.AxureGroup{},
		&models.Axure{},
		&models.Attachment{})
}

func Seed() {
	acct.DBSeed()
}
