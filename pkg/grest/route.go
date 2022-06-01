package grest

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

type Route struct {
	Prefix  string
	Path    string
	Method  string
	Handler func(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}

func RegisterRoute(routes map[string]Route, prefixPath string, as ...*Route) map[string]Route {
	for _, a := range as {
		a.Prefix = prefixPath
		key := a.Prefix + a.Path
		if _, result := routes[key]; result {
			log.Error(routes[key])
			panic("duplicate api definition: " + key)
		}
		routes[key] = *a
	}

	return routes
}
