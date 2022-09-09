package wire_group

import (
	"github.com/google/wire"
	"primedivident/internal/ports/http/asset"
)

var Asset = wire.NewSet(
	asset.NewHandler,
)
