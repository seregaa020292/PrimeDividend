package wire_group

import (
	"github.com/google/wire"
	"primedivident/internal/modules/portfolio/interactor/command"
	"primedivident/internal/modules/portfolio/interactor/query"
	"primedivident/internal/modules/portfolio/repository"

	"primedivident/internal/ports/http/portfolio"
)

var Portfolio = wire.NewSet(
	repository.NewRepository,
	query.NewPortfolioById,
	command.NewPortfolioCreate,
	portfolio.NewHandler,
)
