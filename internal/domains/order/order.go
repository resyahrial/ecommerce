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
	ID                         ksuid.KSUID `json:"ID"`
	BuyerId                    ksuid.KSUID `json:"buyerId"`
	SellerId                   ksuid.KSUID `json:"sellerId"`
	DeliverySourceAddress      string      `json:"deliverySourceAddress"`
	DeliveryDestinationAddress string      `json:"deliveryDestinationAddress"`
	TotalPrice                 float64     `json:"totalPrice"`
	Status                     string      `json:"status"`
	Buyer                      user.Buyer  `json:"buyer"`
	Seller                     user.Seller `json:"seller"`
	Items                      []OrderItem `json:"items"`
}

type OrderItem struct {
	ID        ksuid.KSUID     `json:"ID"`
	ProductId ksuid.KSUID     `json:"productId"`
	Quantity  int64           `json:"quantity"`
	Price     float64         `json:"price"`
	Product   product.Product `json:"product"`
}
