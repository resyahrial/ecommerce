package order_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	order_dom_mock "github.com/resyahrial/go-commerce/internal/domains/order/mocks"
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
