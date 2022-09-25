package hash

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordHasher interface {
	Hash() (string, error)
	Verify(hashed string) error
}

type Password string

func (password Password) Hash() (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (password Password) Verify(hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func (password Password) String() string {
	return string(password)
}
