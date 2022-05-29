package order

import (
	"github.com/resyahrial/go-commerce/internal/domains/product"
	"github.com/resyahrial/go-commerce/internal/domains/user"
	"github.com/segmentio/ksuid"
)

type Order struct {
	ID                         ksuid.KSUID `json:"ID"`
	Buyer                      user.Buyer  `json:"buyer"`
	Seller                     user.Seller `json:"seller"`
	DeliverySourceAddress      string      `json:"deliverySourceAddress"`
	DeliveryDestinationAddress string      `json:"deliveryDestinationAddress"`
	Items                      []OrderItem `json:"items"`
	TotalPrice                 float64     `json:"totalPrice"`
	Status                     string      `json:"status"`
}

type OrderItem struct {
	ID       ksuid.KSUID     `json:"ID"`
	Product  product.Product `json:"product"`
	Quantity int64           `json:"quantity"`
	Price    float64         `json:"price"`
}
