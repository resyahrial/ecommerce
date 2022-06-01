package product

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	product_dom "github.com/resyahrial/go-commerce/internal/domains/product"
	user_dom "github.com/resyahrial/go-commerce/internal/domains/user"
	product_ucase "github.com/resyahrial/go-commerce/internal/usecases/product"
	"github.com/resyahrial/go-commerce/pkg/gctx"
	"github.com/resyahrial/go-commerce/pkg/grest"
	"github.com/resyahrial/go-commerce/pkg/gtrace"
)

type ProductHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	ViewList(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}

type ProductHandler struct {
	productUcase product_ucase.ProductUsecaseInterface
}

func New(productUcase product_ucase.ProductUsecaseInterface) ProductHandlerInterface {
	return &ProductHandler{productUcase}
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var err error
	var prod product_dom.Product

	newCtx, span := gtrace.Start(r.Context())
	defer gtrace.End(span, err)

	if err = grest.ReadRequestBody(r, &prod); err != nil {
		panic(err)
	}

	if prod, err = h.productUcase.Create(newCtx, prod); err != nil {
		panic(err)
	}

	grest.WriteResponse(w, grest.Response{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   prod,
	})
}

func (h *ProductHandler) ViewList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var err error
	var params product_ucase.GetListParams
	var products []product_dom.Product
	var count int64
	var actor gctx.Actor

	newCtx, span := gtrace.Start(r.Context())
	defer gtrace.End(span, err)

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

	if actor, _ = gctx.GetActor(newCtx); actor.Role == user_dom.SELLER {
		params.SellerId = actor.ID
	}

	if products, count, err = h.productUcase.GetList(newCtx, params); err != nil {
		panic(err)
	}

	grest.WriteResponse(w, grest.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data: map[string]interface{}{
			"count":    count,
			"products": products,
		},
	})
}
