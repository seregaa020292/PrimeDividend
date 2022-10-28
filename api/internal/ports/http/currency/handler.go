package currency

import (
	"primedividend/api/internal/infrastructure/server/response"
	"primedividend/api/internal/modules/currency/query"
	"primedividend/api/internal/presenters/currency"
)

type HandlerCurrency struct {
	responder    response.Responder
	presenter    currency.Presenter
	queryGetById query.GetById
	queryGetAll  query.GetAll
}

func NewHandler(
	responder response.Responder,
	presenter currency.Presenter,
	queryGetById query.GetById,
	queryGetAll query.GetAll,
) HandlerCurrency {
	return HandlerCurrency{
		responder:    responder,
		presenter:    presenter,
		queryGetById: queryGetById,
		queryGetAll:  queryGetAll,
	}
}
