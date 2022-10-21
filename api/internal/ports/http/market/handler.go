package market

import (
	"primedivident/internal/infrastructure/server/response"
	"primedivident/internal/modules/market/query"
	"primedivident/internal/presenters/market"
)

type HandlerMarket struct {
	responder    response.Responder
	presenter    market.Presenter
	queryGetById query.GetById
	queryGetAll  query.GetAll
}

func NewHandler(
	responder response.Responder,
	presenter market.Presenter,
	queryGetById query.GetById,
	queryGetAll query.GetAll,
) HandlerMarket {
	return HandlerMarket{
		responder:    responder,
		presenter:    presenter,
		queryGetById: queryGetById,
		queryGetAll:  queryGetAll,
	}
}
