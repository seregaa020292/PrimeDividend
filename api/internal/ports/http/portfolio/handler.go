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
	queryPortfolioAll  query.PortfolioAll
	cmdPortfolioCreate command.PortfolioCreate
	cmdPortfolioEdit   command.PortfolioEdit
	cmdPortfolioRemove command.PortfolioRemove
}

func NewHandler(
	responder response.Responder,
	presenter portfolio.Presenter,
	queryPortfolioById query.PortfolioById,
	queryPortfolioAll query.PortfolioAll,
	cmdPortfolioCreate command.PortfolioCreate,
	cmdPortfolioEdit command.PortfolioEdit,
	cmdPortfolioRemove command.PortfolioRemove,
) HandlerPortfolio {
	return HandlerPortfolio{
		responder:          responder,
		presenter:          presenter,
		queryPortfolioById: queryPortfolioById,
		queryPortfolioAll:  queryPortfolioAll,
		cmdPortfolioCreate: cmdPortfolioCreate,
		cmdPortfolioEdit:   cmdPortfolioEdit,
		cmdPortfolioRemove: cmdPortfolioRemove,
	}
}
