package product

import (
	"context"

	"github.com/mitchellh/mapstructure"
	product_dom "github.com/resyahrial/go-commerce/internal/domains/product"
	"github.com/resyahrial/go-commerce/internal/exceptions"
	"github.com/resyahrial/go-commerce/pkg/gtrace"
	"github.com/segmentio/ksuid"
)

type ProductUsecaseInterface interface {
	GetList(ctx context.Context, params GetListParams) ([]product_dom.Product, int64, error)
	Create(ctx context.Context, product product_dom.Product) (product_dom.Product, error)
}

type GetListParams struct {
	Page     int
	Limit    int
	SellerId ksuid.KSUID
}

func (p GetListParams) ToRepoParams() (repoParams product_dom.GetListParams, err error) {
	if err = mapstructure.Decode(p, &repoParams); err != nil {
		return
	}

	// seller id existence means, seller want to see their products,
	// so, no need to show them their own detail
	repoParams.PreloadSeller = repoParams.SellerId.IsNil()
	return
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
	if repoParams, err = params.ToRepoParams(); err != nil {
		return
	}

	return u.productRepo.GetList(newCtx, repoParams)
}

func (u *ProductUsecase) Create(ctx context.Context, product product_dom.Product) (res product_dom.Product, err error) {
	newCtx, span := gtrace.Start(ctx)
	defer gtrace.End(span, err)

	if errDesc, ok := product.Validate(); !ok {
		err = exceptions.ProductInvalidInputValidation.New(errDesc)
		return
	}

	return u.productRepo.Create(newCtx, product)
}
