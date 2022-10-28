package categorize

import (
	"primedividend/api/internal/modules/auth/service/strategy/auth"
)

type (
	PasswordStrategies = maps[PasswordStrategy]
	PasswordStrategy   interface {
		Login(identify, password string, accountability auth.Accountability) (auth.Tokens, error)
	}
)
