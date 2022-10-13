package wire_group

import (
	"github.com/google/wire"

	"primedivident/internal/config"
	"primedivident/internal/modules/market/service/quotes"
	http "primedivident/internal/ports/http/market"
	ws "primedivident/internal/ports/ws/market"
)

func ProvideQuotes(config config.Config) *quotes.Quotes {
	tinkoff := quotes.NewTinkoff(config.Tinkoff)

	return quotes.NewQuotes(tinkoff)
}

var Market = wire.NewSet(
	ProvideQuotes,
	ws.NewHandlerMarket,
	http.NewHandler,
)
