package wire_group

import (
	"github.com/google/wire"
	"primedivident/internal/ports/http/register"
)

var Register = wire.NewSet(
	register.NewHandler,
)
