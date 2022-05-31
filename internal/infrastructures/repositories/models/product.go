package models

import (
	"time"

	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          ksuid.KSUID `json:"id" gorm:"primaryKey;type:varchar(30)"`
	SellerId    ksuid.KSUID `json:"sellerId" gorm:"index;type:varchar(30)"`
	InsertedAt  time.Time   `json:"insertedAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time   `json:"updatedAt" gorm:"autoUpdateTime"`
	IsDeleted   bool        `json:"isDeleted"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Price       float64     `json:"price"`
	User        User        `json:"-" gorm:"foreignKey:SellerId;references:ID;OnDelete:SET NULL" mapstructure:"-" validate:"-"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID.IsNil() {
		p.ID = ksuid.New()
	}
	return
}
