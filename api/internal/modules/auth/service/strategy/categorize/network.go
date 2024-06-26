package categorize

import (
	"primedividend/api/internal/modules/auth/service/strategy/auth"
)

type (
	NetworkStrategies = maps[NetworkStrategy]
	NetworkStrategy   interface {
		Callback(state string) string
		Login(code string, accountability auth.Accountability) (auth.Tokens, error)
	}
)
