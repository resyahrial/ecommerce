package api_v1

import (
	authentication "github.com/resyahrial/go-commerce/internal/interfaces/http/api/v1/authentications"
	"github.com/resyahrial/go-commerce/pkg/grest"
)

const (
	Prefix = "/api/v1"
)

func GetRoute() map[string]grest.Route {
	Routes := make(map[string]grest.Route)

	authentication.Register(Routes)

	return Routes
}
