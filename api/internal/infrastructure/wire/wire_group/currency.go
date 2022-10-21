package wire_group

import (
	"github.com/google/wire"

	"primedivident/internal/modules/currency/query"
	"primedivident/internal/modules/currency/repository"
	http "primedivident/internal/ports/http/currency"
	presenter "primedivident/internal/presenters/currency"
)

var Currency = wire.NewSet(
	presenter.NewPresenter,
	repository.NewRepository,
	query.NewGetById,
	query.NewGetAll,
	http.NewHandler,
)
