package wire_group

import (
	"github.com/google/wire"

	"primedividend/api/internal/modules/currency/query"
	"primedividend/api/internal/modules/currency/repository"
	http "primedividend/api/internal/ports/http/currency"
	presenter "primedividend/api/internal/presenters/currency"
)

var Currency = wire.NewSet(
	presenter.NewPresenter,
	repository.NewRepository,
	query.NewGetById,
	query.NewGetAll,
	http.NewHandler,
)
