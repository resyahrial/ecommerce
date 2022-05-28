package product

//go:generate mockgen -destination=mocks/mock.go -source=repo.go ProductRepo

type ProductRepo interface{}
