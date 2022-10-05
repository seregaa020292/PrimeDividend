package wire_group

import (
	"github.com/google/wire"

	http "primedivident/internal/ports/http/market"
	ws "primedivident/internal/ports/ws/market"
)

var Market = wire.NewSet(
	ws.NewHandlerMarket,
	http.NewHandler,
)
