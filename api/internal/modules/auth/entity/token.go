package entity

import (
	"time"

	"github.com/google/uuid"

	"primedivident/pkg/datetime"
)

type Token struct {
	Value   uuid.UUID
	Expires time.Time
}

func NewToken(d time.Duration) Token {
	return Token{
		Value:   uuid.New(),
		Expires: datetime.GetNow().Add(d),
	}
}

func (t Token) String() string {
	return t.Value.String()
}

func (t Token) GetExpires() time.Time {
	return t.Expires
}

func (t Token) Valid(value string, expires time.Time) bool {
	return t.isEqual(value) && !t.isExpired(expires)
}

func (t Token) isEqual(value string) bool {
	return t.String() == value
}

func (t Token) isExpired(expires time.Time) bool {
	return t.Expires.After(expires)
}
