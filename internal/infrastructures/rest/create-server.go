package rest

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/resyahrial/go-commerce/config/app"
	"github.com/resyahrial/go-commerce/pkg/grest"
)

func CreateServer(prefix string, routes map[string]grest.Route) {
	router := httprouter.New()

	for path, route := range routes {
		router.Handle(route.Method, prefix+path, route.Handler)
	}

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", app.Host, app.Port),
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
