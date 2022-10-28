package wire_group

import (
	"github.com/google/wire"

	"primedividend/api/internal/modules/portfolio/command"
	"primedividend/api/internal/modules/portfolio/query"
	"primedividend/api/internal/modules/portfolio/repository"
	http "primedividend/api/internal/ports/http/portfolio"
	presenter "primedividend/api/internal/presenters/portfolio"
)

var Portfolio = wire.NewSet(
	presenter.NewPresenter,
	repository.NewRepository,
	query.NewGetById,
	query.NewGetAll,
	query.NewGetUserAll,
	command.NewCreate,
	command.NewEdit,
	command.NewRemove,
	http.NewHandler,
)
