package models

import (
	"time"

	"github.com/segmentio/ksuid"
	"github.com/thoas/go-funk"
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
	Buyer                      User        `json:"buyer" gorm:"foreignKey:BuyerId;references:ID;OnDelete:SET NULL"  validate:"-"`
	Seller                     User        `json:"seller" gorm:"foreignKey:SellerId;references:ID;OnDelete:SET NULL" validate:"-"`
	Items                      []OrderItem `json:"items" gorm:"foreignKey:OrderId;references:ID;OnDelete:SET NULL" validate:"-"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	if o.ID.IsNil() {
		o.ID = ksuid.New()
	}
	return
}

func (o *Order) AfterFind(tx *gorm.DB) (err error) {
	o.TotalPrice = funk.Reduce(o.Items, func(acc float64, item OrderItem) float64 {
		return acc + (item.Price * float64(item.Quantity))
	}, 0).(float64)
	return
}
