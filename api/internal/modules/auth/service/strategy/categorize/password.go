package categorize

import (
	"primedivident/internal/modules/auth/service/strategy/auth"
)

type (
	PasswordStrategies = maps[PasswordStrategy]
	PasswordStrategy   interface {
		Login(identify, password string) (auth.Tokens, error)
	}
)
