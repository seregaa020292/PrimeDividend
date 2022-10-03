package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Session struct {
	ID      uuid.UUID
	Token   string
	Expires time.Time
}

func (s Session) IsExpired(expires time.Time) bool {
	return s.Expires.Before(expires)
}

func (s Session) IsExpiredByNow() bool {
	return s.IsExpired(time.Now())
}

func (s Session) ErrorIsExpiredByNow() error {
	if s.IsExpiredByNow() {
		return errors.New("session in expires")
	}
	return nil
}
