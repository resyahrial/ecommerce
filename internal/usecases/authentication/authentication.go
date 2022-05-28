package authentication

import (
	"context"

	"github.com/resyahrial/go-commerce/internal/domains/authentication"
	"github.com/resyahrial/go-commerce/internal/domains/user"
	"github.com/resyahrial/go-commerce/internal/exception"
	"github.com/resyahrial/go-commerce/pkg/gtrace"
	"github.com/resyahrial/go-commerce/pkg/inspect"
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
	newCtx, span := gtrace.Start(ctx)
	defer gtrace.End(span, err)

	if errDesc, ok := input.Validate(); !ok {
		err = exception.AuthInvalidInput.New(errDesc)
		return
	}

	user := user.User{Email: input.Email}
	if user, err = u.userRepo.GetDetail(newCtx, user); err != nil {
		return
	}

	inspect.Do(user)

	return
}
