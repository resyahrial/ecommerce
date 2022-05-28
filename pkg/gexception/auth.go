package gexception

import "net/http"

func (e *E) InitAuthErr() {
	RegisterException(
		// module name
		"Auth",

		// 4XX
		AuthNotAuthorized,
		AuthForbidden,
		AuthForbiddenRole,
	)
}

var AuthNotAuthorized = &Exception{
	HttpStatus:  http.StatusUnauthorized,
	Code:        "NotAuthorized",
	Description: "need to login to access this",
}

var AuthForbidden = &Exception{
	HttpStatus:  http.StatusForbidden,
	Code:        "Forbidden",
	Description: "already login but not enough permission",
}

var AuthForbiddenRole = &Exception{
	HttpStatus:  http.StatusForbidden,
	Code:        "ForbiddenRole",
	Description: "role not allowed",
}
