package authentication

import (
	"context"

	auth_dom "github.com/resyahrial/go-commerce/internal/domains/authentication"
	"github.com/resyahrial/go-commerce/internal/infrastructures/repositories/models"
	"github.com/resyahrial/go-commerce/pkg/gtrace"
	"gorm.io/gorm"
)

type AuthenticationRepoPg struct {
	db *gorm.DB
}

func New(db *gorm.DB) auth_dom.AuthenticationRepo {
	return &AuthenticationRepoPg{db}
}

func (r *AuthenticationRepoPg) Create(ctx context.Context, refreshToken string) (err error) {
	newCtx, span := gtrace.Start(ctx)
	defer gtrace.End(span, err)

	return r.db.WithContext(newCtx).Create(&models.Authentication{Token: refreshToken}).Error
}
