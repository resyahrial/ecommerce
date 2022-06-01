package order

import (
	"github.com/resyahrial/go-commerce/internal/domains/product"
	"github.com/resyahrial/go-commerce/internal/domains/user"
	"github.com/resyahrial/go-commerce/pkg/gvalidator"
	"github.com/segmentio/ksuid"
)

const (
	PENDING  = "PENDING"
	ACCEPTED = "ACCEPTED"
)

type Order struct {
	ID                         ksuid.KSUID `json:"id"`
	BuyerId                    ksuid.KSUID `json:"buyerId"`
	SellerId                   ksuid.KSUID `json:"sellerId"`
	DeliverySourceAddress      string      `json:"deliverySourceAddress" mapstructure:",omitempty"`
	DeliveryDestinationAddress string      `json:"deliveryDestinationAddress" mapstructure:",omitempty"`
	TotalPrice                 float64     `json:"totalPrice" mapstructure:"-"`
	Status                     string      `json:"status" mapstructure:",omitempty" validate:"omitempty,oneof=PENDING ACCEPTED"`
	Buyer                      user.Buyer  `json:"buyer" mapstructure:"-" validate:"-"`
	Seller                     user.Seller `json:"seller" mapstructure:"-" validate:"-"`
	Items                      []OrderItem `json:"items" mapstructure:",omitempty" validate:"-"`
}

func (o Order) Validate() (string, bool) {
	return gvalidator.Validate(o)
}

type OrderItem struct {
	ID        ksuid.KSUID     `json:"id"`
	ProductId ksuid.KSUID     `json:"productId"`
	Quantity  int64           `json:"quantity" mapstructure:",omitempty"`
	Price     float64         `json:"price" mapstructure:",omitempty"`
	Product   product.Product `json:"product" mapstructure:"-"`
}
