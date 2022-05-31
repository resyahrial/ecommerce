package models

import (
	"time"

	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type Order struct {
	ID                         ksuid.KSUID `json:"ID" gorm:"primaryKey;type:varchar(30)"`
	BuyerId                    ksuid.KSUID `json:"buyerId" gorm:"index;type:varchar(30)"`
	SellerId                   ksuid.KSUID `json:"sellerId" gorm:"index;type:varchar(30)"`
	InsertedAt                 time.Time   `json:"insertedAt" gorm:"autoCreateTime"`
	UpdatedAt                  time.Time   `json:"updatedAt" gorm:"autoUpdateTime"`
	IsDeleted                  bool        `json:"isDeleted"`
	DeliverySourceAddress      string      `json:"deliverySourceAddress"`
	DeliveryDestinationAddress string      `json:"deliveryDestinationAddress"`
	TotalPrice                 float64     `json:"totalPrice"`
	Status                     string      `json:"status"`
	Buyer                      User        `json:"buyer" gorm:"foreignKey:UserId;references:BuyerId;OnDelete:SET NULL"  validate:"-"`
	Seller                     User        `json:"seller" gorm:"foreignKey:UserId;references:SellerId;OnDelete:SET NULL" validate:"-"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	if o.ID.IsNil() {
		o.ID = ksuid.New()
	}
	return
}
