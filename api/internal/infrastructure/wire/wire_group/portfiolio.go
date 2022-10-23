package wire_group

import (
	"github.com/google/wire"

	"primedivident/internal/modules/portfolio/command"
	"primedivident/internal/modules/portfolio/query"
	"primedivident/internal/modules/portfolio/repository"
	http "primedivident/internal/ports/http/portfolio"
	presenter "primedivident/internal/presenters/portfolio"
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
