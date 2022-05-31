package models

import (
	"gorm.io/gorm"
)

func AutoMigrateAllTables(db *gorm.DB) {
	if err := db.AutoMigrate(
		User{},
		Authentication{},
		Product{},
		Order{},
		OrderItem{},
	); err != nil {
		panic(err)
	}
}
