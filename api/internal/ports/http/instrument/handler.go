package instrument

import (
	"net/http"
	"primedivident/internal/infrastructure/server/http/response"
	"primedivident/internal/modules/instrument/repository"
	"primedivident/pkg/logger"
)

type handler struct {
	logger     logger.Logger
	repository repository.Repository
}

func NewHandler(
	logger logger.Logger,
	repository repository.Repository,
) ServerInterface {
	return handler{
		logger:     logger,
		repository: repository,
	}
}

func (h handler) GetInstruments(w http.ResponseWriter, r *http.Request) {
	respond := response.New(w, r)

	instruments, err := h.repository.GetAll()
	if err != nil {
		respond.Err(err)
		return
	}

	respond.Json(http.StatusOK, presenterGetAll(instruments))
}
