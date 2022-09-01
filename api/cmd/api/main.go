package main

import (
	"primedivident/internal/config"
	"primedivident/internal/config/consts"
	"primedivident/internal/infrastructures/server/http/handler"
	"primedivident/internal/infrastructures/wire"
	"primedivident/pkg/graceful"
	"primedivident/pkg/logger"
)

func main() {
	cfg := config.GetConfig()

	logger.SetConfig(logger.Config{
		Format:  consts.TimestampFormat,
		FileLog: consts.TmpLog,
		Level:   cfg.App.LogLevel,
	})

	server := wire.Initialize()

	g := graceful.NewGraceful(consts.TimeoutShutdown)
	g.Shutdown(graceful.Operations{
		server.Stop,
	})

	server.Run(handler.Handlers)
}
