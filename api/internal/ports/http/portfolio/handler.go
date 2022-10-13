package portfolio

import (
	"primedivident/internal/infrastructure/server/response"
	"primedivident/internal/modules/portfolio/command"
	"primedivident/internal/modules/portfolio/query"
	"primedivident/internal/presenters/portfolio"
)

type HandlerPortfolio struct {
	responder          response.Responder
	presenter          portfolio.Presenter
	queryPortfolioById query.PortfolioById
	cmdPortfolioCreate command.PortfolioCreate
	cmdPortfolioEdit   command.PortfolioEdit
}

func NewHandler(
	responder response.Responder,
	presenter portfolio.Presenter,
	queryPortfolioById query.PortfolioById,
	cmdPortfolioCreate command.PortfolioCreate,
	cmdPortfolioEdit command.PortfolioEdit,
) HandlerPortfolio {
	return HandlerPortfolio{
		responder:          responder,
		presenter:          presenter,
		queryPortfolioById: queryPortfolioById,
		cmdPortfolioCreate: cmdPortfolioCreate,
		cmdPortfolioEdit:   cmdPortfolioEdit,
	}
}
