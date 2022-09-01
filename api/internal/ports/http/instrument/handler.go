package instrument

import (
	"net/http"
	"primedivident/internal/config"
	"primedivident/internal/infrastructures/server/http/response"
	"primedivident/internal/modules/instrument/entity"
	"primedivident/internal/modules/instrument/repository"
	"primedivident/pkg/db/postgres"
)

type handler struct {
	repository repository.Repository
}

func NewHandler() ServerInterface {
	return handler{
		repository: repository.NewRepository(postgres.NewPostgres(config.GetConfig().Postgres)),
	}
}

func (h handler) GetInstruments(w http.ResponseWriter, r *http.Request) {
	httpResponse := response.New(w, r)

	instruments, err := h.repository.GetAll()
	if err != nil {
		httpResponse.Err(err)
		return
	}

	httpResponse.Json(http.StatusOK, presenter(instruments))
}

func presenter(instruments entity.Instruments) Instruments {
	result := make(Instruments, len(instruments))

	for i, instrument := range instruments {
		result[i] = Instrument{
			Id:          instrument.ID,
			Title:       instrument.Title,
			Description: instrument.Description,
			CreatedAt:   instrument.CreatedAt,
			UpdatedAt:   instrument.UpdatedAt,
		}
	}

	return result
}
