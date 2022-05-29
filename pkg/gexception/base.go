package gexception

import "net/http"

func (e *E) InitBaseErr() {
	RegisterException(
		// module name
		"Base",

		// 4XX
		BaseInternalServerError,
	)
}

var BaseInternalServerError = &Exception{
	HttpStatus:  http.StatusInternalServerError,
	Code:        "InternalServerError",
	Description: "Internal server error",
}
