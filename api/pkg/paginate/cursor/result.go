package cursor

import (
	"reflect"
	"time"

	"github.com/google/uuid"

	"primedivident/pkg/utils/gog"
)

const (
	KeyID   = "ID"
	KeyTime = "CreatedAt"
)

type PaginateResult[T any] struct {
	Records    []T
	Length     int
	Limit      int
	CursorNext *string
	CursorPrev *string
}

func NewPaginateResult[T any](records []T, input PaginateInput) PaginateResult[T] {
	recordLen := len(records)

	result := PaginateResult[T]{
		Records: records,
		Limit:   input.Limit,
		Length:  gog.Min(recordLen, input.Limit),
	}

	if result.isEmpty() {
		return result
	}

	result.cutRecords()

	if input.Cursor.HasAndPrev() {
		gog.Reverse(result.Records)

		result.setCursorNext()

		if input.EqLimitOver(recordLen) {
			result.setCursorPrev()
		}

		return result
	}

	if !input.Cursor.IsEmpty() {
		result.setCursorPrev()
	}

	if input.EqLimitOver(recordLen) {
		result.setCursorNext()
	}

	return result
}

func (p *PaginateResult[T]) isEmpty() bool {
	return p.Length == 0
}

func (p *PaginateResult[T]) cutRecords() {
	p.Records = p.Records[:p.Length]
}

func (p *PaginateResult[T]) setCursorPrev() {
	cursorPrev := NewCursorPrev(cursorField(p.Records[0]))

	if cursor, err := cursorPrev.Encode(); err == nil {
		p.CursorPrev = &cursor
	}
}

func (p *PaginateResult[T]) setCursorNext() {
	cursorNext := NewCursorNext(cursorField(p.Records[p.Length-1]))

	if cursor, err := cursorNext.Encode(); err == nil {
		p.CursorNext = &cursor
	}
}

func cursorField[T any](record T) (uuid.UUID, time.Time) {
	r := reflect.ValueOf(record)

	return r.FieldByName(KeyID).Interface().(uuid.UUID),
		r.FieldByName(KeyTime).Interface().(time.Time)
}
