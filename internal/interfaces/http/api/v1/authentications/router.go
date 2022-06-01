package authentication

import (
	"github.com/resyahrial/go-commerce/config/app"
	"github.com/resyahrial/go-commerce/internal/infrastructures"
	"github.com/resyahrial/go-commerce/pkg/grest"
	tokenmanager "github.com/resyahrial/go-commerce/pkg/token-manager"
)

func init() {
	handler = New(infrastructures.InitAuthenticationUsecase(
		app.DB,
		tokenmanager.JwtTokenManagerOpts{
			KeyAccess:        app.KeyAccess,
			KeyRefresh:       app.KeyRefresh,
			ExpiryAgeAccess:  app.ExpiryAgeAccess,
			ExpiryAgeRefresh: app.ExpiryAgeRefresh,
		},
	))

	grest.RegisterRoute(
		"/authentications",
		LoginApi,
	)
}

var handler AuthenticationHandlerInterface

var LoginApi = &grest.Route{
	Path:    "/login",
	Handler: handler.Login,
}
