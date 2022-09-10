package instrument

import (
	"primedivident/internal/modules/instrument/interactor/query"
)

type HandlerInstrument struct {
	queryInstrumentAll query.InstrumentAll
}

func NewHandler(
	queryInstrumentAll query.InstrumentAll,
) HandlerInstrument {
	return HandlerInstrument{
		queryInstrumentAll: queryInstrumentAll,
	}
}
