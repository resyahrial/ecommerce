package groute

import (
	"net/http"
	"reflect"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

type RouteMap map[string]Route

var Routes RouteMap
var ExceptionType reflect.Type

type Route struct {
	Prefix  string
	Path    string
	Handler func(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}

func RegisterRoute(prefixPath string, as ...*Route) {
	for _, a := range as {
		a.Prefix = prefixPath
		key := a.Prefix + a.Path
		if _, result := Routes[key]; result {
			log.Error(Routes[key])
			panic("duplicate api definition: " + key)
		}
		Routes[key] = *a
	}
}
