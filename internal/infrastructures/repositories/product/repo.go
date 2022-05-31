package product

import (
	"context"

	"github.com/mitchellh/mapstructure"
	product_dom "github.com/resyahrial/go-commerce/internal/domains/product"
	"github.com/resyahrial/go-commerce/internal/infrastructures/repositories/models"
	"github.com/resyahrial/go-commerce/pkg/gtrace"
	"github.com/resyahrial/go-commerce/pkg/inspect"
	"gorm.io/gorm"
)

type ProductRepoPg struct {
	db *gorm.DB
}

func New(db *gorm.DB) product_dom.ProductRepo {
	return &ProductRepoPg{db}
}

func (r *ProductRepoPg) GetList(ctx context.Context, params product_dom.GetListParams) (products []product_dom.Product, count int64, err error) {
	newCtx, span := gtrace.Start(ctx)
	defer gtrace.End(span, err)

	inspect.Do(newCtx)

	return
}

func (r *ProductRepoPg) Create(ctx context.Context, input product_dom.Product) (res product_dom.Product, err error) {
	newCtx, span := gtrace.Start(ctx)
	defer gtrace.End(span, err)

	var dataProduct models.Product
	if err = mapstructure.Decode(input, &dataProduct); err != nil {
		return
	}

	if err = r.db.WithContext(newCtx).Create(&dataProduct).Error; err != nil {
		return
	}

	if err = mapstructure.Decode(dataProduct, &res); err != nil {
		return
	}

	return
}
