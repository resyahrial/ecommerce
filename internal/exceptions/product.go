package exceptions

import (
	"net/http"

	"github.com/resyahrial/go-commerce/pkg/gexception"
)

func (e *E) InitProductErr() {
	gexception.RegisterException(
		// module name
		"Product",

		// 4XX
		ProductInvalidInputValidation,
		ProductNotFound,
	)
}

var ProductInvalidInputValidation = &gexception.Exception{
	HttpStatus:  http.StatusBadRequest,
	Code:        "InvalidInputValidation",
	Description: "",
}

var ProductNotFound = &gexception.Exception{
	HttpStatus:  http.StatusNotFound,
	Code:        "NotFound",
	Description: "one or few product not found",
}
