package portfolio

import (
	"primedivident/internal/infrastructure/response"
	"primedivident/internal/modules/portfolio/command"
	"primedivident/internal/modules/portfolio/query"
	"primedivident/internal/presenters/portfolio"
)

type HandlerPortfolio struct {
	responder          response.Responder
	presenter          portfolio.Presenter
	queryPortfolioById query.PortfolioById
	cmdPortfolioCreate command.PortfolioCreate
}

func NewHandler(
	responder response.Responder,
	presenter portfolio.Presenter,
	queryPortfolioById query.PortfolioById,
	cmdPortfolioCreate command.PortfolioCreate,
) HandlerPortfolio {
	return HandlerPortfolio{
		responder:          responder,
		presenter:          presenter,
		queryPortfolioById: queryPortfolioById,
		cmdPortfolioCreate: cmdPortfolioCreate,
	}
}
