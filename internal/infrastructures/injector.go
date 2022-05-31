//go:build wireinject
// +build wireinject

package infrastructures

import (
	"github.com/google/wire"
	auth_repo "github.com/resyahrial/go-commerce/internal/infrastructures/repositories/authentication"
	order_repo "github.com/resyahrial/go-commerce/internal/infrastructures/repositories/order"
	product_repo "github.com/resyahrial/go-commerce/internal/infrastructures/repositories/product"
	user_repo "github.com/resyahrial/go-commerce/internal/infrastructures/repositories/user"
	auth_ucase "github.com/resyahrial/go-commerce/internal/usecases/authentication"
	order_ucase "github.com/resyahrial/go-commerce/internal/usecases/order"
	product_ucase "github.com/resyahrial/go-commerce/internal/usecases/product"
	"github.com/resyahrial/go-commerce/pkg/hasher"
	tokenmanager "github.com/resyahrial/go-commerce/pkg/token-manager"
	"gorm.io/gorm"
)

func InitAuthenticationUsecase(db *gorm.DB, tokenManagerOpts tokenmanager.JwtTokenManagerOpts) auth_ucase.AuthenticationUsecaseInterface {
	wire.Build(
		auth_repo.New,
		hasher.NewBcyptHasher,
		tokenmanager.NewJwtTokenManager,
		user_repo.New,
		auth_ucase.New,
	)
	return nil
}

func InitProductUsecase(db *gorm.DB) product_ucase.ProductUsecaseInterface {
	wire.Build(product_repo.New, product_ucase.New)
	return nil
}

func InitProductUsecase(db *gorm.DB) order_ucase.OrderUsecaseInterface {
	wire.Build(order_repo.New, product_repo.New, order_ucase.New)
	return nil
}
