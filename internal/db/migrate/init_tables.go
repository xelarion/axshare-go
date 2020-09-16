package migrate

import (
	"github.com/xandercheung/acct"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func Migrate() {
	err := acct.MigrateTables()
	if err != nil {
		panic(err)
	}
}

func Seed() {
	_ = acct.MigrateSeeds()
}
