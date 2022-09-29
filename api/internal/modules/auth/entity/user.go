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
	PassHash *string
	Role     Role
	Status   Status
	Token    Token
}

type JwtPayload struct {
	ID   uuid.UUID `json:"id"`
	Role string    `json:"role"`
}

func NewUser(email, name, password string) (User, error) {
	pass, err := hash.Password(password).Hash()
	if err != nil {
		return User{}, err
	}

	return User{
		ID:       uuid.New(),
		Email:    email,
		Name:     name,
		PassHash: &pass,
		Role:     UserRole,
		Status:   WaitStatus,
		Token:    NewTokenTTL(),
	}, nil
}

func NewUserNetwork(email, name string) User {
	return User{
		ID:     uuid.New(),
		Email:  email,
		Name:   name,
		Role:   UserRole,
		Status: ActiveStatus,
	}
}

func (u User) ComparePasswordHash(password string) error {
	return hash.Password(password).Verify(*u.PassHash)
}

func (u User) ErrorIsActiveStatus() error {
	if u.Status.IsActive() {
		return nil
	}

	return errors.New("user is not active")
}

func (u User) JwtPayload() JwtPayload {
	return JwtPayload{
		ID:   u.ID,
		Role: u.Role.String(),
	}
}

func (u User) JwtPayloadValidPassword(password string) (JwtPayload, error) {
	if err := u.ErrorIsActiveStatus(); err != nil {
		return JwtPayload{}, err
	}

	if err := u.ComparePasswordHash(password); err != nil {
		return JwtPayload{}, err
	}

	return u.JwtPayload(), nil
}

func (u User) IsEmpty() bool {
	return u == (User{})
}

func (u User) ErrorIsEmpty() error {
	if u.IsEmpty() {
		return errors.New("user is empty")
	}

	return nil
}
