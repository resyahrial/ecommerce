package order

import (
	"context"

	order_dom "github.com/resyahrial/go-commerce/internal/domains/order"
	product_dom "github.com/resyahrial/go-commerce/internal/domains/product"
	"github.com/resyahrial/go-commerce/pkg/gtrace"
	"github.com/resyahrial/go-commerce/pkg/inspect"
	"github.com/segmentio/ksuid"
)

type OrderUsecaseInterface interface {
	GetList(ctx context.Context, params GetListParams) ([]order_dom.Order, int64, error)
}

type GetListParams struct {
	Page   int
	Limit  int
	UserId ksuid.KSUID
	Role   string
}

func (p GetListParams) ToRepoParams() {

}

type OrderUsecase struct {
	orderRepo   order_dom.OrderRepo
	productRepo product_dom.ProductRepo
}

func New(
	orderRepo order_dom.OrderRepo,
	productRepo product_dom.ProductRepo,
) OrderUsecaseInterface {
	return &OrderUsecase{orderRepo, productRepo}
}

func (u *OrderUsecase) GetList(ctx context.Context, params GetListParams) (orders []order_dom.Order, count int64, err error) {
	newCtx, span := gtrace.Start(ctx)
	defer gtrace.End(span, err)

	inspect.Do(newCtx)
	return
}
