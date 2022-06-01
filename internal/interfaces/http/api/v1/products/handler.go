package product

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProductHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}

type ProductHandler struct {
}

func New() ProductHandlerInterface {
	return &ProductHandler{}
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
