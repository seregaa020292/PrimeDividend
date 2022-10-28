package entity

import (
	"errors"

	"github.com/google/uuid"

	"primedividend/api/internal/models"
	"primedividend/api/pkg/utils/hash"
)

type User struct {
	ID       uuid.UUID
	Name     string
	Email    string
	PassHash *string
	Role     Role
	Status   models.Status
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
		Status:   models.WaitStatus,
		Token:    NewTokenTTL(),
	}, nil
}

func NewUserNetwork(email, name string) User {
	return User{
		ID:     uuid.New(),
		Email:  email,
		Name:   name,
		Role:   UserRole,
		Status: models.ActiveStatus,
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

func (u User) ValidPasswordActive(password string) error {
	if err := u.ErrorIsActiveStatus(); err != nil {
		return err
	}

	if err := u.ComparePasswordHash(password); err != nil {
		return err
	}

	return nil
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

func (u *User) SetGenToken() Token {
	u.Token = NewTokenTTL()

	return u.Token
}
