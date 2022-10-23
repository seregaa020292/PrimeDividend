package asset

import (
	"primedivident/internal/infrastructure/server/response"
	"primedivident/internal/modules/asset/command"
	"primedivident/internal/modules/asset/query"
	"primedivident/internal/presenters/asset"
)

type HandlerAsset struct {
	responder       response.Responder
	presenter       asset.Presenter
	queryGetUserAll query.GetUserAll
	commandCreate   command.Create
	commandEdit     command.Edit
	commandRemove   command.Remove
}

func NewHandler(
	responder response.Responder,
	presenter asset.Presenter,
	queryGetUserAll query.GetUserAll,
	commandCreate command.Create,
	commandEdit command.Edit,
	commandRemove command.Remove,
) HandlerAsset {
	return HandlerAsset{
		responder:       responder,
		presenter:       presenter,
		queryGetUserAll: queryGetUserAll,
		commandCreate:   commandCreate,
		commandEdit:     commandEdit,
		commandRemove:   commandRemove,
	}
}
