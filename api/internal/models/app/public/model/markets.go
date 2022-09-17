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

type Markets struct {
	ID           uuid.UUID `sql:"primary_key"`
	Title        string
	Ticker       string
	Content      *string
	ImageURL     *string
	CurrencyID   uuid.UUID
	InstrumentID uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    *time.Time
}
