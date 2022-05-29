package hasher

import "golang.org/x/crypto/bcrypt"

type BcryptHasher struct{}

func NewBcyptHasher() Hasher {
	return &BcryptHasher{}
}

func (h *BcryptHasher) Compare(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (h *BcryptHasher) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
