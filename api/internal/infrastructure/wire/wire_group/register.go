package wire_group

import (
	"github.com/google/wire"

	"primedividend/api/internal/ports/http/register"
)

var Register = wire.NewSet(
	register.NewHandler,
)
