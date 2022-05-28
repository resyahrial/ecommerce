package authentication

import (
	"context"

	"github.com/resyahrial/go-commerce/internal/domains/authentication"
	"github.com/resyahrial/go-commerce/internal/domains/user"
)

type AuthenticationUsecaseInterface interface {
	Login(ctx context.Context, input authentication.Login) (token authentication.Token, err error)
}

type AuthenticationUsecase struct {
	userRepo user.UserRepo
	authRepo authentication.AuthenticationRepo
}

func New(
	userRepo user.UserRepo,
	authRepo authentication.AuthenticationRepo,
) AuthenticationUsecaseInterface {
	return &AuthenticationUsecase{
		userRepo: userRepo,
		authRepo: authRepo,
	}
}

func (u *AuthenticationUsecase) Login(ctx context.Context, input authentication.Login) (token authentication.Token, err error) {
	return
}
