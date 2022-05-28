package seller

//go:generate mockgen -destination=mocks/mock.go -source=repo.go SellerRepo

type SellerRepo interface{}
