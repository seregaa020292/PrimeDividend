package entity

import (
	"time"

	"github.com/google/uuid"

	"primedivident/internal/config/consts"
	"primedivident/pkg/datetime"
	"primedivident/pkg/utils/gog"
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
	return t.Expires.Before(expires)
}

func (t Token) IsExpiredByNow() bool {
	return t.IsExpired(time.Now())
}
