package wire_group

import (
	"github.com/google/wire"

	"primedivident/internal/config"
	"primedivident/internal/modules/market/query"
	"primedivident/internal/modules/market/repository"
	"primedivident/internal/modules/market/service/quotes"
	http "primedivident/internal/ports/http/market"
	ws "primedivident/internal/ports/ws/market"
	presenter "primedivident/internal/presenters/market"
)

func ProvideQuotes(config config.Config) *quotes.Quotes {
	tinkoff := quotes.NewTinkoff(config.Tinkoff)

	return quotes.NewQuotes(tinkoff)
}

var Market = wire.NewSet(
	ProvideQuotes,
	ws.NewHandlerMarket,
	repository.NewRepository,
	presenter.NewPresenter,
	query.NewGetById,
	query.NewGetByTicker,
	query.NewGetAll,
	http.NewHandler,
)
