package entity

import (
	"errors"

	"github.com/google/uuid"

	"primedivident/pkg/utils/hash"
)

type User struct {
	ID       uuid.UUID
	Name     string
	Email    string
	PassHash string
	Status   Status
	Token    Token
}

type JwtUser struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func NewUser(email, password string) (User, error) {
	pass, err := hash.Password(password).Hash()
	if err != nil {
		return User{}, err
	}

	return User{
		ID:       uuid.New(),
		Email:    email,
		PassHash: pass,
		Status:   Wait,
		Token:    NewTokenTTL(),
	}, nil
}

func (u User) ComparePasswordHash(password string) error {
	return hash.Password(password).Verify(u.PassHash)
}

func (u User) JwtPayload() JwtUser {
	return JwtUser{
		Email: u.Email,
		Name:  u.Name,
	}
}

func (u User) JwtPayloadValidPassword(password string) (JwtUser, error) {
	if !u.Status.IsActive() {
		return JwtUser{}, errors.New("user is not active")
	}

	if err := u.ComparePasswordHash(password); err != nil {
		return JwtUser{}, err
	}

	return u.JwtPayload(), nil
}
