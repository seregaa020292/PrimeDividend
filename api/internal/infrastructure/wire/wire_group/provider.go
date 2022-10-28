package wire_group

import (
	"github.com/google/wire"

	"primedividend/api/internal/modules/provider/query"
	"primedividend/api/internal/modules/provider/repository"
	http "primedividend/api/internal/ports/http/provider"
	presenter "primedividend/api/internal/presenters/provider"
)

var Provider = wire.NewSet(
	presenter.NewPresenter,
	repository.NewRepository,
	query.NewGetById,
	query.NewGetAll,
	http.NewHandler,
)
