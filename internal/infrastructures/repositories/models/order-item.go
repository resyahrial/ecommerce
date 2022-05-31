package models

import (
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type OrderItem struct {
	ID        ksuid.KSUID `json:"id" gorm:"primaryKey;type:varchar(30)"`
	OrderId   ksuid.KSUID `json:"orderId" gorm:"index;type:varchar(30)"`
	ProductId ksuid.KSUID `json:"productId" gorm:"index;type:varchar(30)"`
	Quantity  int64       `json:"quantity"`
	Price     float64     `json:"price"`
	Product   Product     `json:"product" gorm:"foreignKey:ProductId;references:ID;OnDelete:SET NULL" mapstructure:"-" validate:"-"`
}

func (o *OrderItem) BeforeCreate(tx *gorm.DB) (err error) {
	if o.ID.IsNil() {
		o.ID = ksuid.New()
	}
	return
}
