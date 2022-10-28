package portfolio

import (
	"primedividend/api/internal/infrastructure/server/response"
	"primedividend/api/internal/modules/portfolio/command"
	"primedividend/api/internal/modules/portfolio/query"
	"primedividend/api/internal/presenters/portfolio"
)

type HandlerPortfolio struct {
	responder       response.Responder
	presenter       portfolio.Presenter
	queryGetById    query.GetById
	queryGetAll     query.GetAll
	queryGetUserAll query.GetUserAll
	cmdCreate       command.Create
	cmdEdit         command.Edit
	cmdRemove       command.Remove
}

func NewHandler(
	responder response.Responder,
	presenter portfolio.Presenter,
	queryGetById query.GetById,
	queryGetAll query.GetAll,
	queryGetUserAll query.GetUserAll,
	cmdCreate command.Create,
	cmdEdit command.Edit,
	cmdRemove command.Remove,
) HandlerPortfolio {
	return HandlerPortfolio{
		responder:       responder,
		presenter:       presenter,
		queryGetById:    queryGetById,
		queryGetAll:     queryGetAll,
		queryGetUserAll: queryGetUserAll,
		cmdCreate:       cmdCreate,
		cmdEdit:         cmdEdit,
		cmdRemove:       cmdRemove,
	}
}
