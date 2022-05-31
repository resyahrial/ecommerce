package order

import (
	"context"

	order_dom "github.com/resyahrial/go-commerce/internal/domains/order"
	"github.com/resyahrial/go-commerce/pkg/gtrace"
	"github.com/resyahrial/go-commerce/pkg/inspect"
	"github.com/segmentio/ksuid"
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

	inspect.Do(newCtx)

	return
}

func (r *OrderRepoPg) Create(ctx context.Context, input order_dom.Order) (res order_dom.Order, err error) {
	newCtx, span := gtrace.Start(ctx)
	defer gtrace.End(span, err)

	inspect.Do(newCtx)

	return
}

func (r *OrderRepoPg) BulkCreate(ctx context.Context, inputs []order_dom.Order) (res []order_dom.Order, err error) {
	newCtx, span := gtrace.Start(ctx)
	defer gtrace.End(span, err)

	inspect.Do(newCtx)

	return
}

func (r *OrderRepoPg) Update(ctx context.Context, id ksuid.KSUID, input order_dom.Order) (res order_dom.Order, err error) {
	newCtx, span := gtrace.Start(ctx)
	defer gtrace.End(span, err)

	inspect.Do(newCtx)

	return
}
