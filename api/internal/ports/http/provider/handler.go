package provider

import (
	"primedividend/api/internal/infrastructure/server/response"
	"primedividend/api/internal/modules/provider/query"
	"primedividend/api/internal/presenters/provider"
)

type HandlerProvider struct {
	responder    response.Responder
	presenter    provider.Presenter
	queryGetById query.GetById
	queryGetAll  query.GetAll
}

func NewHandler(
	responder response.Responder,
	presenter provider.Presenter,
	queryGetById query.GetById,
	queryGetAll query.GetAll,
) HandlerProvider {
	return HandlerProvider{
		responder:    responder,
		presenter:    presenter,
		queryGetById: queryGetById,
		queryGetAll:  queryGetAll,
	}
}
