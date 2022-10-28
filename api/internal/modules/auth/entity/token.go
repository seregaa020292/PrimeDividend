package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"primedividend/api/internal/config/consts"
	"primedividend/api/pkg/datetime"
	"primedividend/api/pkg/utils/gog"
)

type Token struct {
	Value   *uuid.UUID
	Expires *time.Time
}

func NewToken(d time.Duration) Token {
	return Token{
		Value:   gog.Ptr(uuid.New()),
		Expires: gog.Ptr(datetime.GetNow().Add(d)),
	}
}

func NewTokenTTL() Token {
	return NewToken(consts.TokenJoinTTL)
}

func (t Token) String() string {
	return t.Value.String()
}

func (t Token) GetExpires() time.Time {
	return *t.Expires
}

func (t Token) Valid(value string, expires time.Time) bool {
	return t.IsEqual(value) && !t.IsExpired(expires)
}

func (t Token) IsEqual(value string) bool {
	return t.String() == value
}

func (t Token) IsExpired(expires time.Time) bool {
	if t.Expires == nil {
		return false
	}

	return t.Expires.Before(expires)
}

func (t Token) IsExpiredByNow() bool {
	return t.IsExpired(time.Now())
}

func (t Token) ErrorIsExpiredByNow() error {
	if t.IsExpiredByNow() {
		return errors.New("token expired")
	}
	return nil
}
