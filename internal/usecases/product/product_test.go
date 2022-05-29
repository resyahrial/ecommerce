package product_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	product_dom "github.com/resyahrial/go-commerce/internal/domains/product"
	product_dom_mock "github.com/resyahrial/go-commerce/internal/domains/product/mocks"
	"github.com/resyahrial/go-commerce/internal/domains/user"
	"github.com/resyahrial/go-commerce/internal/usecases/product"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/suite"
)

type productUsecaseSuite struct {
	suite.Suite
	productRepo *product_dom_mock.MockProductRepo
	ucase       product.ProductUsecaseInterface
}

func TestProductUsecase(t *testing.T) {
	suite.Run(t, new(productUsecaseSuite))
}

func (s *productUsecaseSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.productRepo = product_dom_mock.NewMockProductRepo(ctrl)
	s.ucase = product.New(s.productRepo)
}

func (s *productUsecaseSuite) TestParamsToRepoParams_Success() {
	params := product.GetListParams{
		Page:     0,
		Limit:    10,
		SellerId: ksuid.New(),
	}

	repoParams, err := params.ToRepoParams()
	s.Nil(err)
	s.Equal(params.Page, repoParams.Page)
	s.Equal(params.Limit, repoParams.Limit)
	s.Equal(params.SellerId, repoParams.SellerId)
	s.False(repoParams.PreloadSeller)

	noPreloadSellerparams := product.GetListParams{}
	noPreloadRepoParams, err := noPreloadSellerparams.ToRepoParams()
	s.Nil(err)
	s.True(noPreloadRepoParams.PreloadSeller)
}

func (s *productUsecaseSuite) TestGetList_Success() {
	params := product.GetListParams{
		Page:     0,
		Limit:    1,
		SellerId: ksuid.New(),
	}

	repoParams, _ := params.ToRepoParams()

	productList := []product_dom.Product{
		{
			ID: ksuid.New(),
			Seller: user.Seller{
				ID: params.SellerId,
			},
		},
	}
	countRepo := int64(10)

	s.productRepo.EXPECT().GetList(gomock.Any(), repoParams).Return(productList, countRepo, nil)

	products, count, err := s.ucase.GetList(context.Background(), params)
	s.Nil(err)
	s.Equal(countRepo, count)
	s.Equal(productList, products)
}
