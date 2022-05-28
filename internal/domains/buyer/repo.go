package buyer

//go:generate mockgen -destination=mocks/mock.go -source=repo.go BuyerRepo

type BuyerRepo interface{}
