package product

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	product_ucase "github.com/resyahrial/go-commerce/internal/usecases/product"
)

type ProductHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}

type ProductHandler struct {
	productUcase product_ucase.ProductUsecaseInterface
}

func New(productUcase product_ucase.ProductUsecaseInterface) ProductHandlerInterface {
	return &ProductHandler{productUcase}
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
