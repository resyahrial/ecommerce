package exceptions

import (
	"net/http"

	"github.com/resyahrial/go-commerce/pkg/gexception"
)

func (e *E) InitUserErr() {
	gexception.RegisterException(
		// module name
		"User",

		// 4XX
		UserNotFound,
	)
}

var UserNotFound = &gexception.Exception{
	HttpStatus:  http.StatusNotFound,
	Code:        "NotFound",
	Description: "user not found",
}
