package categorize

import (
	"primedivident/internal/modules/auth/entity"
)

type (
	PasswordStrategies = maps[PasswordStrategy]
	PasswordStrategy   interface {
		Login(identify, password string) (entity.Tokens, error)
	}
)
