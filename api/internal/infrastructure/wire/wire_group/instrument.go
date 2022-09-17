package wire_group

import (
	"github.com/google/wire"

	"primedivident/internal/modules/instrument/query"
	"primedivident/internal/modules/instrument/repository"
	port "primedivident/internal/ports/http/instrument"
	presenter "primedivident/internal/presenters/instrument"
)

var Instrument = wire.NewSet(
	presenter.NewPresenter,
	repository.NewRepository,
	query.NewInstrumentAll,
	port.NewHandler,
)
