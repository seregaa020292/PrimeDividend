package instrument

import (
	"primedivident/internal/modules/instrument/query"
	"primedivident/pkg/response"
)

type HandlerInstrument struct {
	responder          response.Responder
	queryInstrumentAll query.InstrumentAll
}

func NewHandler(
	responder response.Responder,
	queryInstrumentAll query.InstrumentAll,
) HandlerInstrument {
	return HandlerInstrument{
		responder:          responder,
		queryInstrumentAll: queryInstrumentAll,
	}
}
