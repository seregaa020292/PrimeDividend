//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"primedivident/internal/config"
	serverHttp "primedivident/internal/infrastructure/server/http"
	"primedivident/internal/infrastructure/server/http/handlers"
	"primedivident/internal/services/email"
)

func Initialize(cfg config.Config) serverHttp.Server {
	wire.Build(
		ProvideLogger,
		ProvidePostgres,
		ProvideMailerObserver,

		email.NewFirstTestSend,

		portfolioSet,
		instrumentSet,

		handlers.NewHandlers,
		serverHttp.NewServer,
	)

	return serverHttp.Server{}
}
