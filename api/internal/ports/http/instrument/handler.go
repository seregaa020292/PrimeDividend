package instrument

import (
	"primedividend/api/internal/infrastructure/server/response"
	"primedividend/api/internal/modules/instrument/query"
	"primedividend/api/internal/presenters/instrument"
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
