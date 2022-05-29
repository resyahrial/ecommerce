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
	)
}

var ProductInvalidInputValidation = &gexception.Exception{
	HttpStatus:  http.StatusBadRequest,
	Code:        "InvalidInputValidation",
	Description: "",
}
