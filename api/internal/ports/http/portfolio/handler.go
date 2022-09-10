package portfolio

import (
	"primedivident/internal/modules/portfolio/interactor/command"
	"primedivident/internal/modules/portfolio/interactor/query"
	"primedivident/pkg/validator"
)

type HandlerPortfolio struct {
	validator          validator.Validator
	queryPortfolioById query.PortfolioById
	cmdPortfolioCreate command.PortfolioCreate
}

func NewHandler(
	validator validator.Validator,
	queryPortfolioById query.PortfolioById,
	cmdPortfolioCreate command.PortfolioCreate,
) HandlerPortfolio {
	return HandlerPortfolio{
		validator:          validator,
		queryPortfolioById: queryPortfolioById,
		cmdPortfolioCreate: cmdPortfolioCreate,
	}
}
