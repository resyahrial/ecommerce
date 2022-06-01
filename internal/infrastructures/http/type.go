package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Path    string
	Handler func(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
