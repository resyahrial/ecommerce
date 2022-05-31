package order_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	order_dom "github.com/resyahrial/go-commerce/internal/domains/order"
	order_dom_mock "github.com/resyahrial/go-commerce/internal/domains/order/mocks"
	product_dom "github.com/resyahrial/go-commerce/internal/domains/product"
	product_dom_mock "github.com/resyahrial/go-commerce/internal/domains/product/mocks"
	user_dom "github.com/resyahrial/go-commerce/internal/domains/user"
	"github.com/resyahrial/go-commerce/internal/usecases/order"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/suite"
)

type orderUsecaseSuite struct {
	suite.Suite
	orderRepo   *order_dom_mock.MockOrderRepo
	productRepo *product_dom_mock.MockProductRepo
	ucase       order.OrderUsecaseInterface
}

func TestOrderUsecase(t *testing.T) {
	suite.Run(t, new(orderUsecaseSuite))
}

func (s *orderUsecaseSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.orderRepo = order_dom_mock.NewMockOrderRepo(ctrl)
	s.productRepo = product_dom_mock.NewMockProductRepo(ctrl)
	s.ucase = order.New(s.orderRepo, s.productRepo)
}

func (s *orderUsecaseSuite) TestParamsToRepoParams_Success() {
	buyerParams := order.GetListParams{
		Page:   0,
		Limit:  10,
		UserId: ksuid.New(),
		Role:   user_dom.BUYER,
	}

	repoParams, err := buyerParams.ToRepoParams()
	s.Nil(err)
	s.Equal(buyerParams.Page, repoParams.Page)
	s.Equal(buyerParams.Limit, repoParams.Limit)
	s.Equal(buyerParams.UserId, repoParams.UserId)
	s.True(repoParams.IsBuyer)

	sellerParams := order.GetListParams{
		Role: user_dom.SELLER,
	}

	repoParams, err = sellerParams.ToRepoParams()
	s.Nil(err)
	s.False(repoParams.IsBuyer)
}

func (s *orderUsecaseSuite) TestParamsValidate_SuccessValidate() {
	validParams := order.GetListParams{
		Page:   0,
		Limit:  10,
		UserId: ksuid.New(),
		Role:   user_dom.BUYER,
	}

	errDesc, ok := validParams.Validate()
	s.True(ok)
	s.Empty(errDesc)

	invalidParams := order.GetListParams{
		Page:  0,
		Limit: 10,
	}

	errDesc, ok = invalidParams.Validate()
	s.False(ok)
	s.NotEmpty(errDesc)

	invalidRoleParams := order.GetListParams{
		Page:   0,
		Limit:  10,
		UserId: ksuid.New(),
		Role:   "ADMIN",
	}

	errDesc, ok = invalidRoleParams.Validate()
	s.False(ok)
	s.NotEmpty(errDesc)
}

func (s *orderUsecaseSuite) TestGetList_Success() {
	params := order.GetListParams{
		Page:   0,
		Limit:  1,
		UserId: ksuid.New(),
		Role:   user_dom.BUYER,
	}
	repoParams, _ := params.ToRepoParams()

	orderList := []order_dom.Order{
		{
			ID:         ksuid.New(),
			Buyer:      user_dom.Buyer{ID: params.UserId},
			Seller:     user_dom.Seller{ID: ksuid.New()},
			Status:     order_dom.PENDING,
			TotalPrice: 1000,
			Items: []order_dom.OrderItem{
				{
					ID: ksuid.New(),
					Product: product_dom.Product{
						ID: ksuid.New(),
					},
					Quantity: 1,
					Price:    1000,
				},
			},
		},
	}
	countFromRepo := int64(10)

	s.orderRepo.EXPECT().GetList(gomock.Any(), repoParams).Return(orderList, countFromRepo, nil)

	orders, count, err := s.ucase.GetList(context.Background(), params)
	s.Nil(err)
	s.Equal(orderList, orders)
	s.Equal(countFromRepo, count)
}

func (s *orderUsecaseSuite) TestCreate_Success() {
	firstSeller := user_dom.Seller{
		ID:      ksuid.New(),
		Address: "JKT",
	}

	secondSeller := user_dom.Seller{
		ID:      ksuid.New(),
		Address: "BDG",
	}

	product1 := product_dom.Product{
		ID:     ksuid.New(),
		Price:  100,
		Seller: secondSeller,
	}

	product2 := product_dom.Product{
		ID:     ksuid.New(),
		Price:  100,
		Seller: firstSeller,
	}

	product3 := product_dom.Product{
		ID:     ksuid.New(),
		Price:  100,
		Seller: firstSeller,
	}

	order := order_dom.Order{
		BuyerId:                    ksuid.New(),
		DeliveryDestinationAddress: "BKS",
		Items: []order_dom.OrderItem{
			{
				ProductId: product1.ID,
				Quantity:  1,
			},
			{
				ProductId: product2.ID,
				Quantity:  2,
			},
			{
				ProductId: product3.ID,
				Quantity:  3,
			},
			{
				ProductId: product1.ID,
				Quantity:  3,
			},
		},
	}
	totalProductSearched := int64(3)

	s.productRepo.EXPECT().GetList(gomock.Any(), product_dom.GetListParams{
		Limit:         int(totalProductSearched),
		Ksuids:        []ksuid.KSUID{product1.ID, product2.ID, product3.ID},
		PreloadSeller: true,
	}).Return([]product_dom.Product{product1, product2, product3}, totalProductSearched, nil)

	s.orderRepo.EXPECT().BulkCreate(gomock.Any(), []order_dom.Order{
		{
			BuyerId:                    order.BuyerId,
			SellerId:                   secondSeller.ID,
			DeliverySourceAddress:      secondSeller.Address,
			DeliveryDestinationAddress: order.DeliveryDestinationAddress,
			Items: []order_dom.OrderItem{
				{
					ProductId: product1.ID,
					Quantity:  4,
					Price:     product1.Price,
				},
			},
		},
		{
			BuyerId:                    order.BuyerId,
			SellerId:                   firstSeller.ID,
			DeliverySourceAddress:      firstSeller.Address,
			DeliveryDestinationAddress: order.DeliveryDestinationAddress,
			Items: []order_dom.OrderItem{
				{
					ProductId: product2.ID,
					Quantity:  2,
					Price:     product2.Price,
				},
				{
					ProductId: product3.ID,
					Quantity:  3,
					Price:     product3.Price,
				},
			},
		},
	}).Return([]order_dom.Order{
		{
			ID:         ksuid.New(),
			Status:     order_dom.PENDING,
			TotalPrice: 500,
		},
		{
			ID:         ksuid.New(),
			Status:     order_dom.PENDING,
			TotalPrice: 400,
		},
	}, nil)

	orders, err := s.ucase.Create(context.Background(), order)
	s.Nil(err)
	s.Len(orders, 2)
}

func (s *orderUsecaseSuite) TestUpdate_Success() {
	orderId := ksuid.New()
	input := order_dom.Order{Status: order_dom.ACCEPTED}

	s.orderRepo.EXPECT().Update(gomock.Any(), orderId, input).Return(
		order_dom.Order{ID: orderId, Status: input.Status},
		nil,
	)

	orderRes, err := s.ucase.Update(context.Background(), orderId, input)
	s.Nil(err)
	s.Equal(orderId, orderRes.ID)
	s.Equal(input.Status, orderRes.Status)
}
