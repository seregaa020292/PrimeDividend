package wire_group

import (
	"github.com/google/wire"

	"primedivident/internal/modules/provider/query"
	"primedivident/internal/modules/provider/repository"
	http "primedivident/internal/ports/http/provider"
	presenter "primedivident/internal/presenters/provider"
)

var Provider = wire.NewSet(
	presenter.NewPresenter,
	repository.NewRepository,
	query.NewGetById,
	query.NewGetAll,
	http.NewHandler,
)
