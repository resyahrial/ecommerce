package migrations

import (
	"gorm.io/gorm"
)

type migrationFn func(db *gorm.DB) error

var migrationList []migrationFn = []migrationFn{
	seedUser,
}

func AutoMigration(db *gorm.DB) {
	var err error
	for _, migration := range migrationList {
		if err = migration(db); err != nil {
			panic(err)
		}
	}
}
