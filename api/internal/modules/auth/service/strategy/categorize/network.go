package categorize

import (
	"primedivident/internal/modules/auth/service/strategy/auth"
)

type (
	NetworkStrategies = maps[NetworkStrategy]
	NetworkStrategy   interface {
		Callback(state string) string
		Login(code string) (auth.Tokens, error)
	}
)
