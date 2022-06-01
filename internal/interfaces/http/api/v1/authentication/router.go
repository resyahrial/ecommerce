package authentication

import (
	"github.com/resyahrial/go-commerce/config/app"
	"github.com/resyahrial/go-commerce/internal/infrastructures"
	"github.com/resyahrial/go-commerce/internal/infrastructures/http"
	tokenmanager "github.com/resyahrial/go-commerce/pkg/token-manager"
)

var handler AuthenticationHandlerInterface

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
}

var AuthenticationApi []http.Route = []http.Route{
	{
		Path:    "/login",
		Handler: handler.Login,
	},
}
