package order

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	order_dom "github.com/resyahrial/go-commerce/internal/domains/order"
	order_ucase "github.com/resyahrial/go-commerce/internal/usecases/order"
	"github.com/resyahrial/go-commerce/pkg/gctx"
	"github.com/resyahrial/go-commerce/pkg/grest"
	"github.com/resyahrial/go-commerce/pkg/gtrace"
	"github.com/segmentio/ksuid"
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
	var err error
	var input order_dom.Order
	var orders []order_dom.Order

	newCtx, span := gtrace.Start(r.Context())
	defer gtrace.End(span, err)

	if err = grest.ReadRequestBody(r, &input); err != nil {
		panic(err)
	}

	actor, _ := gctx.GetActor(newCtx)
	input.BuyerId = actor.ID

	if orders, err = h.orderUcase.Create(newCtx, input); err != nil {
		panic(err)
	}

	grest.WriteResponse(w, grest.Response{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   orders,
	})
}

func (h *OrderHandler) ViewList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var err error
	var params order_ucase.GetListParams
	var orders []order_dom.Order
	var count int64

	newCtx, span := gtrace.Start(r.Context())
	defer gtrace.End(span, err)

	actor, _ := gctx.GetActor(newCtx)
	params.UserId = actor.ID
	params.Role = actor.Role

	queries := r.URL.Query()
	if page := queries.Get("page"); page != "" {
		if params.Page, err = strconv.Atoi(page); err != nil {
			panic(err)
		}
	}

	if limit := queries.Get("limit"); limit != "" {
		if params.Limit, err = strconv.Atoi(limit); err != nil {
			panic(err)
		}
	}

	if orders, count, err = h.orderUcase.GetList(newCtx, params); err != nil {
		panic(err)
	}

	grest.WriteResponse(w, grest.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data: map[string]interface{}{
			"count":  count,
			"orders": orders,
		},
	})
}

func (h *OrderHandler) Accept(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var err error
	var od order_dom.Order
	var orderId ksuid.KSUID

	newCtx, span := gtrace.Start(r.Context())
	defer gtrace.End(span, err)

	if err = orderId.Scan(p.ByName("id")); err != nil {
		panic(err)
	}

	if err = grest.ReadRequestBody(r, &od); err != nil {
		panic(err)
	}

	if od, err = h.orderUcase.Update(newCtx, orderId, od); err != nil {
		panic(err)
	}

	grest.WriteResponse(w, grest.Response{
		Code:   http.StatusOK,
		Status: "Updated",
		Data:   od,
	})
}
