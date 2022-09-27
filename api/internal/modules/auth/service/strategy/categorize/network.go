package categorize

import (
	"primedivident/internal/modules/auth/entity"
)

type (
	NetworkStrategies = maps[NetworkStrategy]
	NetworkStrategy   interface {
		Callback(state string) string
		Login(code string) (entity.Tokens, error)
	}
)
