package asset

import (
	"primedivident/internal/infrastructure/server/response"
	"primedivident/internal/modules/asset/query"
	"primedivident/internal/presenters/asset"
)

type HandlerAsset struct {
	responder       response.Responder
	presenter       asset.Presenter
	queryGetUserAll query.GetUserAll
}

func NewHandler(
	responder response.Responder,
	presenter asset.Presenter,
	queryGetUserAll query.GetUserAll,
) HandlerAsset {
	return HandlerAsset{
		responder:       responder,
		presenter:       presenter,
		queryGetUserAll: queryGetUserAll,
	}
}
