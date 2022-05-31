package models

import (
	"time"

	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type Authentication struct {
	ID         ksuid.KSUID `json:"ID" gorm:"primaryKey;type:varchar(30)" `
	InsertedAt time.Time   `json:"insertedAt" gorm:"autoCreateTime"`
	UpdatedAt  time.Time   `json:"updatedAt" gorm:"autoUpdateTime"`
	IsDeleted  bool        `json:"isDeleted"`
	Token      string      `json:"token"`
}

func (a *Authentication) BeforeCreate(tx *gorm.DB) (err error) {
	if a.ID.IsNil() {
		a.ID = ksuid.New()
	}
	return
}
