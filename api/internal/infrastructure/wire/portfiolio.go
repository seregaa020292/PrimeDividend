package wire

import (
	"github.com/google/wire"
	"primedivident/internal/modules/portfolio/interactor/query"
	"primedivident/internal/modules/portfolio/repository"

	"primedivident/internal/ports/http/portfolio"
)

var portfolioSet = wire.NewSet(
	repository.NewRepository,
	query.NewPortfolioById,
	portfolio.NewHandler,
)
