package order

import (
	"github.com/resyahrial/go-commerce/internal/domains/product"
	"github.com/resyahrial/go-commerce/internal/domains/user"
	"github.com/segmentio/ksuid"
)

const (
	PENDING  = "PENDING"
	ACCEPTED = "ACCEPTED"
)

type Order struct {
	ID                         ksuid.KSUID `json:"id" mapstructure:"-"`
	BuyerId                    ksuid.KSUID `json:"buyerId" mapstructure:"-"`
	SellerId                   ksuid.KSUID `json:"sellerId" mapstructure:"-"`
	DeliverySourceAddress      string      `json:"deliverySourceAddress" mapstructure:",omitempty"`
	DeliveryDestinationAddress string      `json:"deliveryDestinationAddress" mapstructure:",omitempty"`
	TotalPrice                 float64     `json:"totalPrice" mapstructure:"-"`
	Status                     string      `json:"status" mapstructure:",omitempty"`
	Buyer                      user.Buyer  `json:"buyer" mapstructure:"-"`
	Seller                     user.Seller `json:"seller" mapstructure:"-"`
	Items                      []OrderItem `json:"items" mapstructure:",omitempty"`
}

type OrderItem struct {
	ID        ksuid.KSUID     `json:"id" mapstructure:"-"`
	ProductId ksuid.KSUID     `json:"productId" mapstructure:"-"`
	Quantity  int64           `json:"quantity" mapstructure:",omitempty"`
	Price     float64         `json:"price" mapstructure:",omitempty"`
	Product   product.Product `json:"product" mapstructure:"-"`
}
