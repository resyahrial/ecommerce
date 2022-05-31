package product

import (
	"context"

	"github.com/mitchellh/mapstructure"
	product_dom "github.com/resyahrial/go-commerce/internal/domains/product"
	user_dom "github.com/resyahrial/go-commerce/internal/domains/user"
	"github.com/resyahrial/go-commerce/internal/infrastructures/repositories/models"
	"github.com/resyahrial/go-commerce/pkg/gtrace"
	"github.com/thoas/go-funk"
	"gorm.io/gorm"
)

type ProductRepoPg struct {
	db *gorm.DB
}

func New(db *gorm.DB) product_dom.ProductRepo {
	return &ProductRepoPg{db}
}

func (r *ProductRepoPg) GetList(ctx context.Context, params product_dom.GetListParams) (res []product_dom.Product, count int64, err error) {
	newCtx, span := gtrace.Start(ctx)
	defer gtrace.End(span, err)

	var products []models.Product

	offset := params.Page * params.Limit
	result := r.db.WithContext(newCtx)
	queries := []string{}
	values := []interface{}{}

	if params.Ksuids != nil {
		queries = append(queries, "ksuid IN ?")
		values = append(values, params.Ksuids)
	}

	if !params.SellerId.IsNil() {
		queries = append(queries, "seller_id = ?")
		values = append(values, params.SellerId)
	}

	if params.PreloadSeller {
		result = result.Preload("User")
	}

	joinedQuery := funk.Reduce(queries, func(acc string, query string) string {
		return acc + query
	}, "")

	result = result.Model(&models.Product{}).Where(joinedQuery, values...)

	if err = result.Count(&count).Error; err != nil {
		return
	}

	result = result.Limit(params.Limit).Offset(offset)

	if err = result.Find(&products).Error; err != nil {
		return
	}

	for _, p := range products {
		var productDomain product_dom.Product
		if productDomain, err = r.castDataModelsToDomain(newCtx, p); err != nil {
			return
		}
		res = append(res, productDomain)
	}

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

func (r *ProductRepoPg) castDataModelsToDomain(ctx context.Context, p models.Product) (pDom product_dom.Product, err error) {
	_, span := gtrace.Start(ctx)
	defer gtrace.End(span, err)

	var userDomain user_dom.User
	if err = mapstructure.Decode(p, &pDom); err != nil {
		return
	}

	if err = mapstructure.Decode(p.User, &userDomain); err != nil {
		return
	}

	if seller, ok := userDomain.ToSeller(); ok {
		pDom.Seller = seller
	}

	return
}
