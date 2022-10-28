package asset

import (
	"primedividend/api/internal/infrastructure/server/response"
	"primedividend/api/internal/modules/asset/command"
	"primedividend/api/internal/modules/asset/query"
	"primedividend/api/internal/presenters/asset"
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
