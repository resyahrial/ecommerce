package product

import (
	"github.com/resyahrial/go-commerce/internal/domains/user"
	"github.com/segmentio/ksuid"
)

type Product struct {
	ID          ksuid.KSUID `json:"ID"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Price       float64     `json:"price"`
	Seller      user.Seller `json:"seller"`
}
