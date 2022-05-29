package order

import (
	"context"

	"github.com/mitchellh/mapstructure"
	order_dom "github.com/resyahrial/go-commerce/internal/domains/order"
	product_dom "github.com/resyahrial/go-commerce/internal/domains/product"
	user_dom "github.com/resyahrial/go-commerce/internal/domains/user"
	"github.com/resyahrial/go-commerce/internal/exceptions"
	"github.com/resyahrial/go-commerce/pkg/gtrace"
	"github.com/resyahrial/go-commerce/pkg/gvalidator"
	"github.com/segmentio/ksuid"
)

type OrderUsecaseInterface interface {
	GetList(ctx context.Context, params GetListParams) ([]order_dom.Order, int64, error)
	Create(ctx context.Context, buyerId ksuid.KSUID, order order_dom.Order) ([]order_dom.Order, error)
}

type GetListParams struct {
	Page   int
	Limit  int
	UserId ksuid.KSUID `validate:"required"`
	Role   string      `validate:"required,oneof=BUYER SELLER"`
}

func (p GetListParams) Validate() (string, bool) {
	return gvalidator.Validate(p)
}

func (p GetListParams) ToRepoParams() (repoParams order_dom.GetListParams, err error) {
	if err = mapstructure.Decode(p, &repoParams); err != nil {
		return
	}

	repoParams.IsBuyer = p.Role == user_dom.BUYER
	return
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

	var repoParams order_dom.GetListParams

	if errDesc, ok := params.Validate(); !ok {
		err = exceptions.OrderInvalidInputValidation.New(errDesc)
		return
	}

	if repoParams, err = params.ToRepoParams(); err != nil {
		return
	}

	return u.orderRepo.GetList(newCtx, repoParams)
}

func (u *OrderUsecase) Create(ctx context.Context, buyerId ksuid.KSUID, order order_dom.Order) (orders []order_dom.Order, err error) {
	newCtx, span := gtrace.Start(ctx)
	defer gtrace.Error(span, err)

	var productKsuids []ksuid.KSUID
	var products []product_dom.Product
	var productCount int64
	sellerOrderItemMap := make(map[ksuid.KSUID][]order_dom.OrderItem)
	ksuidQuantityMap := make(map[ksuid.KSUID]int64)

	for _, orderItem := range order.Items {
		productKsuids = append(productKsuids, orderItem.ProductId)
		ksuidQuantityMap[orderItem.ProductId] = orderItem.Quantity
	}

	if products, productCount, err = u.productRepo.GetList(newCtx, product_dom.GetListParams{
		Limit:         len(productKsuids),
		Ksuids:        productKsuids,
		PreloadSeller: true,
	}); err != nil {
		return
	} else if int(productCount) != len(productKsuids) {
		err = exceptions.ProductNotFound
		return
	}

	for _, product := range products {
		sellerOrderItemMap[product.Seller.ID] = append(
			sellerOrderItemMap[product.Seller.ID],
			order_dom.OrderItem{
				ProductId: product.ID,
				Quantity:  ksuidQuantityMap[product.ID],
				Price:     product.Price,
			},
		)
	}

	for sellerId, items := range sellerOrderItemMap {
		orders = append(orders, order_dom.Order{
			BuyerId:                    buyerId,
			SellerId:                   sellerId,
			DeliverySourceAddress:      order.DeliverySourceAddress,
			DeliveryDestinationAddress: order.DeliveryDestinationAddress,
			Status:                     order_dom.PENDING,
			Items:                      items,
		})
	}

	return u.orderRepo.BulkCreate(newCtx, orders)
}
