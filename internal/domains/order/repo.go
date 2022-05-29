package order

import (
	"context"

	"github.com/segmentio/ksuid"
)

//go:generate mockgen -destination=mocks/mock.go -source=repo.go OrderRepo

type OrderRepo interface {
	GetList(ctx context.Context, params GetListParams) ([]Order, int64, error)
	Create(ctx context.Context, input Order) (Order, error)
	BulkCreate(ctx context.Context, inputs []Order) ([]Order, error)
}

type GetListParams struct {
	Page    int
	Limit   int
	UserId  ksuid.KSUID
	IsBuyer bool
}
