package authentication

import (
	"net/http"

	"github.com/resyahrial/go-commerce/config/app"
	"github.com/resyahrial/go-commerce/internal/infrastructures"
	"github.com/resyahrial/go-commerce/pkg/grest"
	tokenmanager "github.com/resyahrial/go-commerce/pkg/token-manager"
)

func Register(routes map[string]grest.Route) {
	authUcase := infrastructures.InitAuthenticationUsecase(
		app.DB,
		tokenmanager.JwtTokenManagerOpts{
			KeyAccess:        app.KeyAccess,
			KeyRefresh:       app.KeyRefresh,
			ExpiryAgeAccess:  app.ExpiryAgeAccess,
			ExpiryAgeRefresh: app.ExpiryAgeRefresh,
		},
	)
	handler := New(authUcase)

	grest.RegisterRoute(
		routes,
		"/authentications",
		&grest.Route{
			Path:    "/login",
			Method:  http.MethodPost,
			Handler: handler.Login,
		},
	)
}
