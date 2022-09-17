package wire_group

import (
	"github.com/google/wire"

	"primedivident/internal/modules/portfolio/command"
	"primedivident/internal/modules/portfolio/query"
	"primedivident/internal/modules/portfolio/repository"
	port "primedivident/internal/ports/http/portfolio"
	presenter "primedivident/internal/presenter/portfolio"
)

var Portfolio = wire.NewSet(
	presenter.NewPresenter,
	repository.NewRepository,
	query.NewPortfolioById,
	command.NewPortfolioCreate,
	port.NewHandler,
)
