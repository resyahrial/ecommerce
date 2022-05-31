package models

import (
	"gorm.io/gorm"
)

func AutoMigrateAllTables(db *gorm.DB) {
	if err := db.AutoMigrate(
		User{},
	); err != nil {
		panic(err)
	}
}
