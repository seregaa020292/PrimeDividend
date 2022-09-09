package wire_group

import (
	"github.com/google/wire"
	"primedivident/internal/ports/http/user"
)

var User = wire.NewSet(
	user.NewHandler,
)
