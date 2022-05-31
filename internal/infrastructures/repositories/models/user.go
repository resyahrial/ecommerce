package models

import (
	"time"

	"github.com/segmentio/ksuid"
)

type User struct {
	ID         ksuid.KSUID `json:"ID" gorm:"primaryKey;type:varchar(30)"`
	InsertedAt time.Time   `json:"insertedAt" gorm:"autoCreateTime"`
	UpdatedAt  time.Time   `json:"updatedAt" gorm:"autoUpdateTime"`
	IsDeleted  bool        `json:"isDeleted"`
	Email      string      `json:"email"`
	Name       string      `json:"name"`
	Password   string      `json:"-"`
	Address    string      `json:"address"`
	Role       string      `json:"-"`
}
