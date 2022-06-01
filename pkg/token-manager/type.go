package tokenmanager

//go:generate mockgen -destination=mocks/mock.go -source=type.go TokenManager

type TokenManager interface {
	GenerateAccess(claims Claims) (string, bool)
	GenerateRefresh(claims Claims) (string, bool)
	ParseAccess(tokenString string) (Claims, error)
	ParseRefresh(tokenString string) (Claims, error)
}

type Claims struct {
	ID string `json:"id"`
}
