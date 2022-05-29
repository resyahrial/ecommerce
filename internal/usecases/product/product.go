package product

import (
	"context"

	"github.com/mitchellh/mapstructure"
	product_dom "github.com/resyahrial/go-commerce/internal/domains/product"
	"github.com/resyahrial/go-commerce/pkg/gtrace"
	"github.com/segmentio/ksuid"
)

type ProductUsecaseInterface interface {
	GetList(ctx context.Context, params GetListParams) ([]product_dom.Product, int64, error)
}

type GetListParams struct {
	Page     int
	Limit    int
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

func (u *ProductUsecase) GetList(ctx context.Context, params GetListParams) (products []product_dom.Product, count int64, err error) {
	newCtx, span := gtrace.Start(ctx)
	defer gtrace.End(span, err)

	var repoParams product_dom.GetListParams
	if err = mapstructure.Decode(params, &repoParams); err != nil {
		return
	}

	return u.productRepo.GetList(newCtx, repoParams)
}
