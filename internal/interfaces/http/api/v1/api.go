package api_v1

import (
	authentication "github.com/resyahrial/go-commerce/internal/interfaces/http/api/v1/authentications"
	"github.com/resyahrial/go-commerce/pkg/grest"
)

var Routes map[string]grest.Route

const (
	Prefix = "/api/v1"
)

func init() {
	Routes = make(map[string]grest.Route)

	authentication.Register(Routes)
}
