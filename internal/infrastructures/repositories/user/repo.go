package user

import (
	"context"

	"github.com/mitchellh/mapstructure"
	user_dom "github.com/resyahrial/go-commerce/internal/domains/user"
	"github.com/resyahrial/go-commerce/internal/exceptions"
	"github.com/resyahrial/go-commerce/internal/infrastructures/repositories/models"
	"github.com/resyahrial/go-commerce/pkg/gtrace"
	"gorm.io/gorm"
)

type UserRepoPg struct {
	db *gorm.DB
}

func New(db *gorm.DB) user_dom.UserRepo {
	return &UserRepoPg{db}
}

func (r *UserRepoPg) GetDetail(ctx context.Context, input user_dom.User) (res user_dom.User, err error) {
	newCtx, span := gtrace.Start(ctx)
	defer gtrace.End(span, err)

	var dataUser models.User
	if err = r.db.WithContext(newCtx).Where(input).First(&dataUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = exceptions.UserNotFound
		}
		return
	}

	if err = mapstructure.Decode(dataUser, &res); err != nil {
		return
	}

	return
}
