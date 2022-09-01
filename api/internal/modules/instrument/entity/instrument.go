package entity

import (
	"time"

	"github.com/google/uuid"
)

type Instrument struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type Instruments = []Instrument
