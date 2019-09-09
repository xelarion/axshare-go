package migrate

import (
	"axshare_go/internal/db"
	"axshare_go/internal/models"
	"fmt"
)

func Migrate() {
	fmt.Println("migrate ...")
	db.AxshareDb.Debug().AutoMigrate(
		&models.User{},
		&models.AxureGroup{},
		&models.Axure{},
		&models.Attachment{})
}
