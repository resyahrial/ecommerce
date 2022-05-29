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
	ID                         ksuid.KSUID `json:"ID" mapstucture:"-"`
	BuyerId                    ksuid.KSUID `json:"buyerId" mapstucture:"-"`
	SellerId                   ksuid.KSUID `json:"sellerId" mapstucture:"-"`
	DeliverySourceAddress      string      `json:"deliverySourceAddress" mapstucture:",omitempty"`
	DeliveryDestinationAddress string      `json:"deliveryDestinationAddress" mapstucture:",omitempty"`
	TotalPrice                 float64     `json:"totalPrice" mapstucture:"-"`
	Status                     string      `json:"status" mapstucture:",omitempty"`
	Buyer                      user.Buyer  `json:"buyer" mapstucture:"-"`
	Seller                     user.Seller `json:"seller" mapstucture:"-"`
	Items                      []OrderItem `json:"items" mapstucture:",omitempty"`
}

type OrderItem struct {
	ID        ksuid.KSUID     `json:"ID" mapstucture:"-"`
	ProductId ksuid.KSUID     `json:"productId" mapstucture:"-"`
	Quantity  int64           `json:"quantity" mapstucture:",omitempty"`
	Price     float64         `json:"price" mapstucture:",omitempty"`
	Product   product.Product `json:"product" mapstucture:"-"`
}
