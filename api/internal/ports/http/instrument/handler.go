package instrument

import (
	"primedivident/internal/infrastructure/server/response"
	"primedivident/internal/modules/instrument/query"
	"primedivident/internal/presenters/instrument"
)

type HandlerInstrument struct {
	responder    response.Responder
	presenter    instrument.Presenter
	queryGetById query.GetById
	queryGetAll  query.GetAll
}

func NewHandler(
	responder response.Responder,
	presenter instrument.Presenter,
	queryGetById query.GetById,
	queryGetAll query.GetAll,
) HandlerInstrument {
	return HandlerInstrument{
		responder:    responder,
		presenter:    presenter,
		queryGetById: queryGetById,
		queryGetAll:  queryGetAll,
	}
}
