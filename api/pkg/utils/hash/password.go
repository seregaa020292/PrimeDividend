package hash

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordHasher interface {
	Hash() (string, error)
	Verify(hashed string) bool
}

type Password string

func (password Password) Hash() (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (password Password) Verify(hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))
	return err == nil
}

func (password Password) String() string {
	return string(password)
}
