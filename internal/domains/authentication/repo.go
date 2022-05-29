package authentication

import "context"

//go:generate mockgen -destination=mocks/mock.go -source=repo.go AuthenticationRepo

type AuthenticationRepo interface {
	Create(ctx context.Context, refreshToken string) error
}
