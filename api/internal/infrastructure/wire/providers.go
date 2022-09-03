package wire

import (
	"primedivident/internal/config"
	"primedivident/internal/config/consts"
	"primedivident/pkg/db/postgres"
	"primedivident/pkg/logger"
)

func ProvidePostgres(cfg config.Config) *postgres.Postgres {
	return postgres.NewPostgres(cfg.Postgres)
}

func ProvideLogger(cfg config.Config) logger.Logger {
	return logger.InitConfig(logger.Config{
		Format:  consts.TimestampFormat,
		FileLog: consts.TmpLog,
		Level:   cfg.App.LogLevel,
	})
}
