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
	Role     Role
	Status   Status
	Token    Token
}

type JwtPayload struct {
	ID   uuid.UUID `json:"id"`
	Role string    `json:"role"`
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
		Role:     UserRole,
		Status:   WaitStatus,
		Token:    NewTokenTTL(),
	}, nil
}

func (u User) ComparePasswordHash(password string) error {
	return hash.Password(password).Verify(u.PassHash)
}

func (u User) JwtPayload() JwtPayload {
	return JwtPayload{
		ID:   u.ID,
		Role: u.Role.String(),
	}
}

func (u User) JwtPayloadValidPassword(password string) (JwtPayload, error) {
	if !u.Status.IsActive() {
		return JwtPayload{}, errors.New("user is not active")
	}

	if err := u.ComparePasswordHash(password); err != nil {
		return JwtPayload{}, err
	}

	return u.JwtPayload(), nil
}
