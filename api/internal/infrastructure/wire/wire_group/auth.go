package wire_group

import (
	"github.com/google/wire"

	"primedivident/internal/ports/http/auth"
)

var Auth = wire.NewSet(
	auth.NewHandler,
)
