//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"primedivident/internal/config"
	serverHttp "primedivident/internal/infrastructure/http"
	"primedivident/internal/infrastructure/http/handlers"
	"primedivident/internal/services/email"
	"primedivident/pkg/validator"
)

func Initialize(cfg config.Config) serverHttp.Server {
	wire.Build(
		ProvideLogger,
		ProvidePostgres,
		ProvideMailerObserver,

		validator.GetValidator,

		email.NewFirstTestSend,

		portfolioSet,
		instrumentSet,

		handlers.NewHandlers,
		serverHttp.NewServer,
	)

	return serverHttp.Server{}
}
