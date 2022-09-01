//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	serverHttp "primedivident/internal/infrastructures/server/http"
)

func Initialize() serverHttp.Server {
	wire.Build(
		serverHttp.NewServer,
	)

	return serverHttp.Server{}
}
