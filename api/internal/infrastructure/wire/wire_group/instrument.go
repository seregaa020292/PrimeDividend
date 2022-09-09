package wire_group

import (
	"github.com/google/wire"

	"primedivident/internal/modules/instrument/interactor/query"
	"primedivident/internal/modules/instrument/repository"
	"primedivident/internal/ports/http/instrument"
)

var Instrument = wire.NewSet(
	repository.NewRepository,
	query.NewInstrumentAll,
	instrument.NewHandler,
)
