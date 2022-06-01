package order

import (
	"net/http"

	"github.com/resyahrial/go-commerce/config/app"
	user_dom "github.com/resyahrial/go-commerce/internal/domains/user"
	"github.com/resyahrial/go-commerce/internal/infrastructures"
	"github.com/resyahrial/go-commerce/pkg/grest"
)

func Register(routes map[string]grest.Route) {
	orderUcase := infrastructures.InitOrderUsecase(app.DB)
	handler := New(orderUcase)

	grest.RegisterRoute(
		routes,
		"/orders",
		&grest.Route{
			Path:       "/",
			Method:     http.MethodPost,
			Handler:    handler.Create,
			IsNeedAuth: true,
			Role:       user_dom.BUYER,
		},
		&grest.Route{
			Path:       "/",
			Method:     http.MethodGet,
			Handler:    handler.ViewList,
			IsNeedAuth: true,
		},
		&grest.Route{
			Path:       "/accept",
			Method:     http.MethodPatch,
			Handler:    handler.Accept,
			IsNeedAuth: true,
			Role:       user_dom.SELLER,
		},
	)
}
