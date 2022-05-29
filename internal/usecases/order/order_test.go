package order_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	order_dom "github.com/resyahrial/go-commerce/internal/domains/order"
	order_dom_mock "github.com/resyahrial/go-commerce/internal/domains/order/mocks"
	"github.com/resyahrial/go-commerce/internal/domains/product"
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

func TestProductUsecase(t *testing.T) {
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
					Product: product.Product{
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
