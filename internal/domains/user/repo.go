package user

//go:generate mockgen -destination=mocks/mock.go -source=repo.go UserRepo

type UserRepo interface{}
