package order

import (
	"context"

	"github.com/segmentio/ksuid"
)

//go:generate mockgen -destination=mocks/mock.go -source=repo.go OrderRepo

type OrderRepo interface {
	GetList(ctx context.Context, params GetListParams) ([]Order, int64, error)
	GetDetailByParams(ctx context.Context, input Order) (Order, error)
	Create(ctx context.Context, input Order) (Order, error)
	BulkCreate(ctx context.Context, inputs []Order) ([]Order, error)
	Update(ctx context.Context, id ksuid.KSUID, input Order) (Order, error)
}

type GetListParams struct {
	Page    int
	Limit   int
	UserId  ksuid.KSUID
	IsBuyer bool
}
