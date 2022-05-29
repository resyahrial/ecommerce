package hasher

//go:generate mockgen -destination=mocks/mock.go -source=type.go Hasher

type Hasher interface {
	Compare(password, hash string) bool
	Hash(password string) (string, error)
}
