package order

import (
	"context"

	"github.com/mitchellh/mapstructure"
	order_dom "github.com/resyahrial/go-commerce/internal/domains/order"
	product_dom "github.com/resyahrial/go-commerce/internal/domains/product"
	user_dom "github.com/resyahrial/go-commerce/internal/domains/user"
	"github.com/resyahrial/go-commerce/internal/infrastructures/repositories/models"
	"github.com/resyahrial/go-commerce/pkg/gtrace"
	"github.com/segmentio/ksuid"
	"github.com/thoas/go-funk"
	"gorm.io/gorm"
)

type OrderRepoPg struct {
	db *gorm.DB
}

func New(db *gorm.DB) order_dom.OrderRepo {
	return &OrderRepoPg{db}
}

func (r *OrderRepoPg) GetList(ctx context.Context, params order_dom.GetListParams) (res []order_dom.Order, count int64, err error) {
	newCtx, span := gtrace.Start(ctx)
	defer gtrace.End(span, err)

	var orders []models.Order

	offset := params.Page * params.Limit
	result := r.db.WithContext(newCtx)
	queries := []string{}
	values := []interface{}{}

	if !params.UserId.IsNil() {
		if params.IsBuyer {
			queries = append(queries, "buyer_id = ?")
		} else {
			queries = append(queries, "seller_id = ?")
		}
		values = append(values, params.UserId)
	}

	result = result.Preload("Buyer").Preload("Seller").Preload("Items", "order_items.is_deleted <> true").Preload("Items.Product")

	joinedQuery := funk.Reduce(queries, func(acc string, query string) string {
		return acc + query
	}, "")

	result = result.Model(&models.Product{}).Where(joinedQuery, values...)

	if err = result.Count(&count).Error; err != nil {
		return
	}

	result = result.Limit(params.Limit).Offset(offset)

	if err = result.Find(&orders).Error; err != nil {
		return
	}

	for _, o := range orders {
		var orderDomain order_dom.Order
		if orderDomain, err = r.castDataModelsToDomain(newCtx, o); err != nil {
			return
		}
		res = append(res, orderDomain)
	}

	return
}

func (r *OrderRepoPg) Create(ctx context.Context, input order_dom.Order) (res order_dom.Order, err error) {
	newCtx, span := gtrace.Start(ctx)
	defer gtrace.End(span, err)

	var dataOrder models.Order
	if err = mapstructure.Decode(input, &dataOrder); err != nil {
		return
	}

	if err = r.db.WithContext(newCtx).Create(&dataOrder).Error; err != nil {
		return
	}

	if res, err = r.castDataModelsToDomain(newCtx, dataOrder); err != nil {
		return
	}

	return
}

func (r *OrderRepoPg) BulkCreate(ctx context.Context, inputs []order_dom.Order) (res []order_dom.Order, err error) {
	newCtx, span := gtrace.Start(ctx)
	defer gtrace.End(span, err)

	var dataOrders []models.Order
	if err = mapstructure.Decode(inputs, &dataOrders); err != nil {
		return
	}

	if err = r.db.WithContext(newCtx).Create(&dataOrders).Error; err != nil {
		return
	}

	for _, o := range dataOrders {
		var orderDomain order_dom.Order
		if orderDomain, err = r.castDataModelsToDomain(newCtx, o); err != nil {
			return
		}
		res = append(res, orderDomain)
	}

	return
}

func (r *OrderRepoPg) Update(ctx context.Context, id ksuid.KSUID, input order_dom.Order) (res order_dom.Order, err error) {
	newCtx, span := gtrace.Start(ctx)
	defer gtrace.End(span, err)

	var dataOrder models.Order
	if err = mapstructure.Decode(input, &dataOrder); err != nil {
		return
	}

	if err = r.db.WithContext(newCtx).Model(&models.Order{}).Where("ksuid = ?", id).Omit("Buyer", "Seller", "Items").Updates(&dataOrder).Error; err != nil {
		return
	}

	if res, err = r.castDataModelsToDomain(newCtx, dataOrder); err != nil {
		return
	}

	return
}

func (r *OrderRepoPg) castDataModelsToDomain(ctx context.Context, o models.Order) (oDom order_dom.Order, err error) {
	_, span := gtrace.Start(ctx)
	defer gtrace.End(span, err)

	var buyerDom user_dom.User
	var sellerDom user_dom.User
	var orderItems []order_dom.OrderItem
	if err = mapstructure.Decode(o, &oDom); err != nil {
		return
	}

	if err = mapstructure.Decode(o.Buyer, &buyerDom); err != nil {
		return
	} else if buyer, ok := buyerDom.ToBuyer(); ok {
		oDom.Buyer = buyer
	}

	if err = mapstructure.Decode(o.Seller, &sellerDom); err != nil {
		return
	} else if seller, ok := sellerDom.ToSeller(); ok {
		oDom.Seller = seller
	}

	for _, item := range o.Items {
		var orderItem order_dom.OrderItem
		var productDom product_dom.Product
		if err = mapstructure.Decode(item, &orderItem); err != nil {
			return
		}

		if err = mapstructure.Decode(item.Product, &productDom); err != nil {
			return
		} else {
			orderItem.Product = productDom
		}
		orderItems = append(orderItems, orderItem)
	}

	oDom.Items = orderItems

	return
}
