package wire_group

import (
	"github.com/google/wire"

	"primedividend/api/internal/modules/instrument/query"
	"primedividend/api/internal/modules/instrument/repository"
	http "primedividend/api/internal/ports/http/instrument"
	presenter "primedividend/api/internal/presenters/instrument"
)

var Instrument = wire.NewSet(
	presenter.NewPresenter,
	repository.NewRepository,
	query.NewGetById,
	query.NewGetAll,
	http.NewHandler,
)
