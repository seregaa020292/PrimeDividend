package wire_group

import (
	"github.com/google/wire"

	"primedivident/internal/infrastructure/wire/providers"
	"primedivident/internal/modules/market/query"
	"primedivident/internal/modules/market/repository"
	http "primedivident/internal/ports/http/market"
	ws "primedivident/internal/ports/ws/market"
	presenter "primedivident/internal/presenters/market"
)

var Market = wire.NewSet(
	providers.ProvideQuotes,
	ws.NewHandlerMarket,
	repository.NewRepository,
	presenter.NewPresenter,
	query.NewGetById,
	query.NewGetByTicker,
	query.NewGetAll,
	http.NewHandler,
)
