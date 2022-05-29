package product

import (
	"github.com/resyahrial/go-commerce/internal/domains/user"
	"github.com/resyahrial/go-commerce/pkg/gvalidator"
	"github.com/segmentio/ksuid"
)

type Product struct {
	ID          ksuid.KSUID `json:"ID"`
	Name        string      `json:"name" validate:"required,max=50"`
	Description string      `json:"description" validate:"required"`
	Price       float64     `json:"price" validate:"required,gte=0"`
	Seller      user.Seller `json:"seller"`
}

func (p Product) Validate() (string, bool) {
	return gvalidator.Validate(p)
}
