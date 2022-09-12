package portfolio

import (
	"primedivident/internal/modules/portfolio/interactor/command"
	"primedivident/internal/modules/portfolio/interactor/query"
	"primedivident/pkg/response"
)

type HandlerPortfolio struct {
	responder          response.Responder
	queryPortfolioById query.PortfolioById
	cmdPortfolioCreate command.PortfolioCreate
}

func NewHandler(
	responder response.Responder,
	queryPortfolioById query.PortfolioById,
	cmdPortfolioCreate command.PortfolioCreate,
) HandlerPortfolio {
	return HandlerPortfolio{
		responder:          responder,
		queryPortfolioById: queryPortfolioById,
		cmdPortfolioCreate: cmdPortfolioCreate,
	}
}
