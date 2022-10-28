package cursor

import (
	"time"

	"github.com/google/uuid"

	"primedividend/api/pkg/utils/gog"
)

type Cursor struct {
	ID        uuid.UUID
	Time      time.Time
	Direction Direction
}

func NewCursor(id uuid.UUID, time time.Time, direction Direction) Cursor {
	return Cursor{
		ID:        id,
		Time:      time,
		Direction: direction,
	}
}

func NewCursorNext(id uuid.UUID, time time.Time) Cursor {
	return NewCursor(id, time, NextDirection)
}

func NewCursorPrev(id uuid.UUID, time time.Time) Cursor {
	return NewCursor(id, time, PrevDirection)
}

func (c Cursor) IsEmpty() bool {
	return c == (Cursor{})
}

func (c Cursor) HasAndPrev() bool {
	return !c.IsEmpty() && c.Direction.IsPrev()
}

func (c Cursor) HasAndNext() bool {
	return !c.IsEmpty() && c.Direction.IsNext()
}

func (c Cursor) Encode() (string, error) {
	return Encode(c)
}

func (c Cursor) Operator() string {
	return gog.If(c.HasAndPrev(), ">", "<")
}
