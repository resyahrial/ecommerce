package exception

import (
	"net/http"

	"github.com/resyahrial/go-commerce/pkg/gexception"
)

func (e *E) InitAuthErr() {
	gexception.RegisterException(
		// module name
		"Auth",

		// 4XX
		AuthInvalidInput,
		AuthNotAuthorized,
		AuthForbidden,
		AuthForbiddenRole,
	)
}

var AuthInvalidInput = &gexception.Exception{
	HttpStatus:  http.StatusBadRequest,
	Code:        "InvalidInput",
	Description: "",
}

var AuthNotAuthorized = &gexception.Exception{
	HttpStatus:  http.StatusUnauthorized,
	Code:        "NotAuthorized",
	Description: "need to login to access this",
}

var AuthForbidden = &gexception.Exception{
	HttpStatus:  http.StatusForbidden,
	Code:        "Forbidden",
	Description: "already login but not enough permission",
}

var AuthForbiddenRole = &gexception.Exception{
	HttpStatus:  http.StatusForbidden,
	Code:        "ForbiddenRole",
	Description: "role not allowed",
}
