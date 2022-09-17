package instrument

import (
	"primedivident/internal/modules/instrument/query"
	"primedivident/internal/presenters/instrument"
	"primedivident/pkg/response"
)

type HandlerInstrument struct {
	responder          response.Responder
	presenter          instrument.Presenter
	queryInstrumentAll query.InstrumentAll
}

func NewHandler(
	responder response.Responder,
	presenter instrument.Presenter,
	queryInstrumentAll query.InstrumentAll,
) HandlerInstrument {
	return HandlerInstrument{
		responder:          responder,
		presenter:          presenter,
		queryInstrumentAll: queryInstrumentAll,
	}
}
