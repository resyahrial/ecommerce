package product

import (
	"context"

	"github.com/segmentio/ksuid"
)

//go:generate mockgen -destination=mocks/mock.go -source=repo.go ProductRepo

type ProductRepo interface {
	GetList(ctx context.Context, params GetListParams) ([]Product, error)
}

type GetListParams struct {
	Page     int
	Limit    int
	SellerId ksuid.KSUID
}
