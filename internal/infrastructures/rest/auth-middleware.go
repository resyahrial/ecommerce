package rest

import (
	"context"
	"net/http"

	"github.com/mitchellh/mapstructure"
	"github.com/resyahrial/go-commerce/config/app"
	"github.com/resyahrial/go-commerce/internal/domains/user"
	"github.com/resyahrial/go-commerce/internal/exceptions"
	user_repo "github.com/resyahrial/go-commerce/internal/infrastructures/repositories/user"
	"github.com/resyahrial/go-commerce/pkg/gctx"
	"github.com/resyahrial/go-commerce/pkg/gtrace"
	tokenmanager "github.com/resyahrial/go-commerce/pkg/token-manager"
	"github.com/segmentio/ksuid"
	log "github.com/sirupsen/logrus"
)

type AuthMiddleware struct {
	tokenManager tokenmanager.TokenManager
	userRepo     user.UserRepo
	nextHandler  http.Handler
}

func NewAuthMiddleware(nextHandler http.Handler) *AuthMiddleware {
	tokenManager := tokenmanager.NewJwtTokenManager(
		tokenmanager.JwtTokenManagerOpts{
			KeyAccess:        app.KeyAccess,
			KeyRefresh:       app.KeyRefresh,
			ExpiryAgeAccess:  app.ExpiryAgeAccess,
			ExpiryAgeRefresh: app.ExpiryAgeRefresh,
		},
	)

	userRepo := user_repo.New(app.DB)
	return &AuthMiddleware{tokenManager: tokenManager, userRepo: userRepo, nextHandler: nextHandler}
}

func (m *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	var userLogin user.User
	var actor gctx.Actor

	newCtx, span := gtrace.Start(r.Context())
	defer gtrace.End(span, err)

	token := r.Header.Get("authorization")
	if userLogin, err = m.tokenValidation(newCtx, token); err != nil {
		panic(err)
	}

	if err = mapstructure.Decode(userLogin, &actor); err != nil {
		panic(err)
	}

	newCtx = gctx.SetDataAndGetNewCtx(newCtx, gctx.CtxData{Actor: actor})

	m.nextHandler.ServeHTTP(w, r.WithContext(newCtx))
}

func (m *AuthMiddleware) tokenValidation(ctx context.Context, token string) (userLogin user.User, err error) {
	newCtx, span := gtrace.Start(ctx)
	defer gtrace.End(span, err)

	var claims tokenmanager.Claims
	var userId ksuid.KSUID

	if claims, err = m.tokenManager.ParseAccess(token); err != nil {
		log.Error(err)
		err = exceptions.AuthNotAuthorized
		return
	}

	if err = userId.Scan(claims.ID); err != nil {
		return
	}

	if userLogin, err = m.userRepo.GetDetail(newCtx, user.User{ID: userId}); err != nil {
		err = exceptions.AuthFailed
		return
	}

	return
}
