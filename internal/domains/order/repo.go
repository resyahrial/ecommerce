package order

//go:generate mockgen -destination=mocks/mock.go -source=repo.go OrderRepo

type OrderRepo interface{}
