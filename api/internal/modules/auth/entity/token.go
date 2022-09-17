package entity

import (
	"primedivident/pkg/utils/gog"
	"time"

	"github.com/google/uuid"

	"primedivident/pkg/datetime"
)

type Token struct {
	Value   *string   `db:"token_join_value"`
	Expires time.Time `db:"token_join_expires"`
}

func NewToken(d time.Duration) Token {
	return Token{
		Value:   gog.Ptr(uuid.New().String()),
		Expires: datetime.GetNow().Add(d),
	}
}

func (t Token) String() string {
	return *t.Value
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
