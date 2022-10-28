package wire_group

import (
	"github.com/google/wire"

	"primedividend/api/internal/infrastructure/wire/providers"
	"primedividend/api/internal/modules/market/query"
	"primedividend/api/internal/modules/market/repository"
	"primedividend/api/internal/modules/market/service/quotes"
	http "primedividend/api/internal/ports/http/market"
	ws "primedividend/api/internal/ports/ws/market"
	presenter "primedividend/api/internal/presenters/market"
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
