package authentication

import (
	"github.com/resyahrial/go-commerce/config/app"
	"github.com/resyahrial/go-commerce/internal/infrastructures"
	api_v1 "github.com/resyahrial/go-commerce/internal/interfaces/http/api/v1"
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

	api_v1.Routes = grest.RegisterRoute(
		api_v1.Routes,
		"/authentications",
		LoginApi,
	)
}

var handler AuthenticationHandlerInterface

var LoginApi = &grest.Route{
	Path:    "/login",
	Handler: handler.Login,
}
