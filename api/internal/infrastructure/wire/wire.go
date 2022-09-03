//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"primedivident/internal/config"
	serverHttp "primedivident/internal/infrastructure/server/http"
	"primedivident/internal/infrastructure/server/http/handlers"
	"primedivident/internal/modules/instrument/repository"
	"primedivident/internal/ports/http/instrument"
	"primedivident/internal/ports/http/portfolio"
)

func Initialize(cfg config.Config) serverHttp.Server {
	wire.Build(
		ProvideLogger,
		ProvidePostgres,

		repository.NewRepository,

		portfolio.NewHandler,
		instrument.NewHandler,

		handlers.NewHandlers,
		serverHttp.NewServer,
	)

	return serverHttp.Server{}
}
