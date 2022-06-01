package api_v1

import (
	auth_handler "github.com/resyahrial/go-commerce/internal/interfaces/http/api/v1/authentications"
	order_handler "github.com/resyahrial/go-commerce/internal/interfaces/http/api/v1/orders"
	product_handler "github.com/resyahrial/go-commerce/internal/interfaces/http/api/v1/products"
	"github.com/resyahrial/go-commerce/pkg/grest"
)

const (
	Prefix = "/api/v1"
)

func GetRoute() map[string]grest.Route {
	Routes := make(map[string]grest.Route)

	auth_handler.Register(Routes)
	product_handler.Register(Routes)
	order_handler.Register(Routes)

	return Routes
}
