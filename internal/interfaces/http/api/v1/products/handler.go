package product

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	product_dom "github.com/resyahrial/go-commerce/internal/domains/product"
	product_ucase "github.com/resyahrial/go-commerce/internal/usecases/product"
	"github.com/resyahrial/go-commerce/pkg/grest"
	"github.com/resyahrial/go-commerce/pkg/gtrace"
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
