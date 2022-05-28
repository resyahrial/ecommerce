package authentication

import (
	"context"

	"github.com/resyahrial/go-commerce/internal/domains/authentication"
)

type AuthenticationUsecaseInterface interface {
	Login(ctx context.Context, input authentication.Login) (token authentication.Token, err error)
}

type AuthenticationUsecase struct {
	authRepo authentication.AuthenticationRepo
}

func New(authRepo authentication.AuthenticationRepo) AuthenticationUsecaseInterface {
	return &AuthenticationUsecase{
		authRepo: authRepo,
	}
}

func (u *AuthenticationUsecase) Login(ctx context.Context, input authentication.Login) (token authentication.Token, err error) {
	return
}
