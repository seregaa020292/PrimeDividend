package entity

import (
	"time"

	"github.com/google/uuid"

	"primedivident/pkg/utils/hash"
)

type User struct {
	ID       uuid.UUID     `db:"id"`
	Name     string        `db:"name"`
	Email    string        `db:"email"`
	Password hash.Password `db:"password"`
	Role     string        `db:"role"`
	Avatar   string        `db:"avatar"`
	Token
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
