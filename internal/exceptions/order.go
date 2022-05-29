package exceptions

import (
	"net/http"

	"github.com/resyahrial/go-commerce/pkg/gexception"
)

func (e *E) InitOrderErr() {
	gexception.RegisterException(
		// module name
		"Order",

		// 4XX
		OrderInvalidInputValidation,
	)
}

var OrderInvalidInputValidation = &gexception.Exception{
	HttpStatus:  http.StatusBadRequest,
	Code:        "InvalidInputValidation",
	Description: "",
}
