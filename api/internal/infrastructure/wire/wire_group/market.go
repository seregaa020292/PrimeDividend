package wire_group

import (
	"github.com/google/wire"

	"primedivident/internal/infrastructure/wire/providers"
	"primedivident/internal/modules/market/query"
	"primedivident/internal/modules/market/repository"
	"primedivident/internal/modules/market/service/quotes"
	http "primedivident/internal/ports/http/market"
	ws "primedivident/internal/ports/ws/market"
	presenter "primedivident/internal/presenters/market"
)

var Market = wire.NewSet(
	providers.ProvideTinkoff,
	ws.NewHandlerMarket,
	repository.NewRepository,
	repository.NewAssetRepository,
	presenter.NewPresenter,
	query.NewGetById,
	query.NewGetByTicker,
	query.NewGetAll,
	quotes.NewHubQuotes,
	http.NewHandler,
)
