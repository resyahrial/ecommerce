package rest

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/resyahrial/go-commerce/config/app"
	api_v1 "github.com/resyahrial/go-commerce/internal/interfaces/http/api/v1"
	"github.com/resyahrial/go-commerce/pkg/gexception"
	"github.com/resyahrial/go-commerce/pkg/grest"
)

func CreateServer() {
	router := httprouter.New()

	for path, route := range api_v1.GetRoute() {
		router.Handle(route.Method, api_v1.Prefix+path, route.Handler)
	}

	router.PanicHandler = panicHandler

	address := fmt.Sprintf("%s:%s", app.Host, app.Port)
	server := http.Server{
		Addr:    address,
		Handler: router,
	}

	fmt.Printf("http://%s\n", address)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func panicHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	var exception *gexception.Exception
	var ok bool
	if exception, ok = err.(*gexception.Exception); !ok {
		exception = gexception.BaseInternalServerError
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(exception.HttpStatus)
	grest.WriteResponse(w, grest.Response{
		Code:   exception.HttpStatus,
		Status: exception.Code,
		Data:   exception.Description,
	})
}
