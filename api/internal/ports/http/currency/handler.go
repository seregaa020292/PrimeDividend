package currency

import (
	"primedivident/internal/infrastructure/server/response"
	"primedivident/internal/modules/currency/query"
	"primedivident/internal/presenters/currency"
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
