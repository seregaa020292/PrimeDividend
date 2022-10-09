package wire_group

import (
	"github.com/google/wire"

	"primedivident/internal/modules/market/service/quotes"
	http "primedivident/internal/ports/http/market"
	ws "primedivident/internal/ports/ws/market"
)

var Market = wire.NewSet(
	quotes.NewQuotes,
	ws.NewHandlerMarket,
	http.NewHandler,
)
