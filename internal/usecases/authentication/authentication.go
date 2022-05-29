package authentication

import (
	"context"

	"github.com/resyahrial/go-commerce/internal/domains/authentication"
	"github.com/resyahrial/go-commerce/internal/domains/user"
	"github.com/resyahrial/go-commerce/internal/exceptions"
	"github.com/resyahrial/go-commerce/pkg/gtrace"
	"github.com/resyahrial/go-commerce/pkg/hasher"
	tokenmanager "github.com/resyahrial/go-commerce/pkg/token-manager"
)

type AuthenticationUsecaseInterface interface {
	Login(ctx context.Context, input authentication.Login) (token authentication.Token, err error)
}

type AuthenticationUsecase struct {
	authRepo     authentication.AuthenticationRepo
	hashHandler  hasher.Hasher
	tokenManager tokenmanager.TokenManager
	userRepo     user.UserRepo
}

func New(
	authRepo authentication.AuthenticationRepo,
	hashHandler hasher.Hasher,
	tokenManager tokenmanager.TokenManager,
	userRepo user.UserRepo,
) AuthenticationUsecaseInterface {
	return &AuthenticationUsecase{authRepo, hashHandler, tokenManager, userRepo}
}

func (u *AuthenticationUsecase) Login(ctx context.Context, input authentication.Login) (token authentication.Token, err error) {
	newCtx, span := gtrace.Start(ctx)
	defer gtrace.End(span, err)

	if errDesc, ok := input.Validate(); !ok {
		err = exceptions.AuthInvalidInputValidation.New(errDesc)
		return
	}

	user := user.User{Email: input.Email}
	if user, err = u.userRepo.GetDetail(newCtx, user); err != nil {
		return
	} else if ok := u.hashHandler.Compare(input.Password, user.Password); !ok || user.ID.IsNil() {
		err = exceptions.AuthInvalidInput
		return
	}

	tokenClaims := tokenmanager.Claims{
		ID: user.ID.String(),
	}

	accessToken, _ := u.tokenManager.GenerateAccess(tokenClaims)
	refreshToken, _ := u.tokenManager.GenerateRefresh(tokenClaims)
	if accessToken == "" || refreshToken == "" {
		err = exceptions.AuthFailed
		return
	}

	token = authentication.Token{
		Access:  accessToken,
		Refresh: refreshToken,
	}

	if err = u.authRepo.Create(newCtx, token.Refresh); err != nil {
		return
	}

	return
}
