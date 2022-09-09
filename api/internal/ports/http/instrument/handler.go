package instrument

import (
	"net/http"
	"primedivident/internal/modules/instrument/interactor/query"
	"primedivident/pkg/logger"
	"primedivident/pkg/response"
)

type HandlerInstrument struct {
	logger             logger.Logger
	queryInstrumentAll query.InstrumentAll
}

func NewHandler(
	logger logger.Logger,
	queryInstrumentAll query.InstrumentAll,
) HandlerInstrument {
	return HandlerInstrument{
		logger:             logger,
		queryInstrumentAll: queryInstrumentAll,
	}
}

func (h HandlerInstrument) GetInstruments(w http.ResponseWriter, r *http.Request) {
	respond := response.New(w, r)

	instruments, err := h.queryInstrumentAll.Fetch(query.FilterOrderInstruments{})
	if err != nil {
		respond.Err(err)
		return
	}

	respond.Json(http.StatusOK, presenterGetAll(instruments))
}
