package entity

import (
	"github.com/google/uuid"
	"time"
)

type Portfolio struct {
	ID         uuid.UUID `db:"id"`
	Title      string    `db:"title"`
	Active     bool      `db:"active"`
	UserID     uuid.UUID `db:"user_id"`
	CurrencyID uuid.UUID `db:"currency_id"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
