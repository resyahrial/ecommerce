package models

import (
	"gorm.io/gorm"
)

func AutoMigrateAllTables(db *gorm.DB) {
	if err := db.AutoMigrate(
		User{},
		Authentication{},
	); err != nil {
		panic(err)
	}
}
