package product

import (
	"context"

	product_dom "github.com/resyahrial/go-commerce/internal/domains/product"
	"github.com/resyahrial/go-commerce/pkg/gtrace"
	"github.com/resyahrial/go-commerce/pkg/inspect"
	"github.com/segmentio/ksuid"
)

type ProductUsecaseInterface interface {
	GetList(ctx context.Context, params GetListParams) ([]product_dom.Product, error)
}

type GetListParams struct {
	Page     int
	Limit    int
	IsMyList bool
	SellerId ksuid.KSUID
}

type ProductUsecase struct {
	productRepo product_dom.ProductRepo
}

func New(
	productRepo product_dom.ProductRepo,
) ProductUsecaseInterface {
	return &ProductUsecase{productRepo}
}

func (u *ProductUsecase) GetList(ctx context.Context, params GetListParams) (products []product_dom.Product, err error) {
	newCtx, span := gtrace.Start(ctx)
	defer gtrace.End(span, err)

	inspect.Do(newCtx)

	return
}
