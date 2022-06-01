package authentication

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	auth_ucase "github.com/resyahrial/go-commerce/internal/usecases/authentication"
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

}
