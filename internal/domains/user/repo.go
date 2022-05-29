package user

import "context"

//go:generate mockgen -destination=mocks/mock.go -source=repo.go UserRepo

type UserRepo interface {
	GetDetail(ctx context.Context, input User) (User, error)
}
