package authentication

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	auth_dom "github.com/resyahrial/go-commerce/internal/domains/authentication"
	auth_ucase "github.com/resyahrial/go-commerce/internal/usecases/authentication"
	"github.com/resyahrial/go-commerce/pkg/grest"
	"github.com/resyahrial/go-commerce/pkg/gtrace"
)

type AuthenticationHandlerInterface interface {
	Login(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}

type AuthenticationHandler struct {
	authUcase auth_ucase.AuthenticationUsecaseInterface
}

func New(authUcase auth_ucase.AuthenticationUsecaseInterface) AuthenticationHandlerInterface {
	return &AuthenticationHandler{authUcase}
}

func (h *AuthenticationHandler) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var err error
	var input auth_dom.Login
	var token auth_dom.Token

	newCtx, span := gtrace.Start(r.Context())
	defer gtrace.End(span, err)

	if err = grest.ReadRequestBody(r, &input); err != nil {
		panic(err)
	}

	if token, err = h.authUcase.Login(newCtx, input); err != nil {
		panic(err)
	}

	grest.WriteResponse(w, grest.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   token,
	})
}
