//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"github.com/google/uuid"
	"time"
)

type Portfolios struct {
	ID         uuid.UUID `sql:"primary_key"`
	Title      string
	Active     bool
	UserID     uuid.UUID
	CurrencyID uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  *time.Time
}
