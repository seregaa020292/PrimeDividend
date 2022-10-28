package user

import (
	"primedividend/api/internal/infrastructure/server/response"
	"primedividend/api/internal/modules/user/command"
	"primedividend/api/internal/modules/user/query"
	"primedividend/api/internal/presenters/user"
)

type HandlerUser struct {
	responder     response.Responder
	presenter     user.Presenter
	queryGetById  query.GetById
	commandRemove command.Remove
	commandEdit   command.Edit
}

func NewHandler(
	responder response.Responder,
	presenter user.Presenter,
	queryGetById query.GetById,
	commandRemove command.Remove,
	commandEdit command.Edit,
) HandlerUser {
	return HandlerUser{
		responder:     responder,
		presenter:     presenter,
		queryGetById:  queryGetById,
		commandRemove: commandRemove,
		commandEdit:   commandEdit,
	}
}
