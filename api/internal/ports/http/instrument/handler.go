package instrument

import (
	"github.com/google/uuid"
	"net/http"
	"primedivident/internal/config"
	"primedivident/internal/infrastructures/server/http/response"
	"primedivident/pkg/db/postgres"
	"time"
)

type handler struct {
}

func NewHandler() ServerInterface {
	return handler{}
}

type ModelInstrument struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (h handler) GetInstruments(w http.ResponseWriter, r *http.Request) {
	pq := postgres.NewPostgres(config.GetConfig().Postgres)

	var instruments []ModelInstrument
	err := pq.Select(&instruments, "SELECT * FROM instruments")
	if err != nil {
		response.New(w, r).Err(err)
		return
	}

	response.New(w, r).Json(http.StatusOK, presenter(instruments))
}

func presenter(instruments []ModelInstrument) Instruments {
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
