package wire_group

import (
	"github.com/google/wire"
	"primedivident/internal/ports/http/currency"
)

var Currency = wire.NewSet(
	currency.NewHandler,
)
