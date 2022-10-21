package provider

import (
	"primedivident/internal/infrastructure/server/response"
	"primedivident/internal/modules/provider/query"
	"primedivident/internal/presenters/provider"
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
