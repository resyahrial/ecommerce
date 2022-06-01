package product

import (
	"net/http"

	"github.com/resyahrial/go-commerce/config/app"
	user_dom "github.com/resyahrial/go-commerce/internal/domains/user"
	"github.com/resyahrial/go-commerce/internal/infrastructures"
	"github.com/resyahrial/go-commerce/pkg/grest"
)

func Register(routes map[string]grest.Route) {
	productUcase := infrastructures.InitProductUsecase(app.DB)
	handler := New(productUcase)

	grest.RegisterRoute(
		routes,
		"/products",
		&grest.Route{
			Path:       "/",
			Method:     http.MethodPost,
			Handler:    handler.Create,
			IsNeedAuth: true,
			Role:       user_dom.SELLER,
		},
	)
}
