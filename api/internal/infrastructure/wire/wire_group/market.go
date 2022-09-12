package wire_group

import (
	"github.com/google/wire"

	"primedivident/internal/ports/http/market"
)

var Market = wire.NewSet(
	market.NewHandler,
)
