package portfolio

import (
	"primedivident/internal/infrastructure/server/response"
	"primedivident/internal/modules/portfolio/command"
	"primedivident/internal/modules/portfolio/query"
	"primedivident/internal/presenters/portfolio"
)

type HandlerPortfolio struct {
	responder    response.Responder
	presenter    portfolio.Presenter
	queryGetById query.GetById
	queryGetAll  query.GetAll
	cmdCreate    command.Create
	cmdEdit      command.Edit
	cmdRemove    command.Remove
}

func NewHandler(
	responder response.Responder,
	presenter portfolio.Presenter,
	queryGetById query.GetById,
	queryGetAll query.GetAll,
	cmdCreate command.Create,
	cmdEdit command.Edit,
	cmdRemove command.Remove,
) HandlerPortfolio {
	return HandlerPortfolio{
		responder:    responder,
		presenter:    presenter,
		queryGetById: queryGetById,
		queryGetAll:  queryGetAll,
		cmdCreate:    cmdCreate,
		cmdEdit:      cmdEdit,
		cmdRemove:    cmdRemove,
	}
}
