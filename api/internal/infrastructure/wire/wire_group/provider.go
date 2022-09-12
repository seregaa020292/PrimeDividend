package wire_group

import (
	"github.com/google/wire"

	"primedivident/internal/ports/http/provider"
)

var Provider = wire.NewSet(
	provider.NewHandler,
)
