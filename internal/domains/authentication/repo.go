package authentication

//go:generate mockgen -destination=mocks/mock.go -source=repo.go AuthenticationRepo

type AuthenticationRepo interface{}
