package order

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	order_ucase "github.com/resyahrial/go-commerce/internal/usecases/order"
)

type OrderHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	ViewList(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Accept(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}

type OrderHandler struct {
	orderUcase order_ucase.OrderUsecaseInterface
}

func New(orderUcase order_ucase.OrderUsecaseInterface) OrderHandlerInterface {
	return &OrderHandler{orderUcase}
}

func (h *OrderHandler) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (h *OrderHandler) ViewList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (h *OrderHandler) Accept(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
