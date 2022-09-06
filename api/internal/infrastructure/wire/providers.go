package wire

import (
	"primedivident/internal/config"
	"primedivident/internal/config/consts"
	"primedivident/pkg/db/postgres"
	"primedivident/pkg/logger"
	"primedivident/pkg/mailer"
)

func ProvideLogger(cfg config.Config) logger.Logger {
	return logger.InitConfig(logger.Config{
		Format:  consts.TimestampFormat,
		FileLog: consts.TmpLog,
		Level:   cfg.App.LogLevel,
	})
}

func ProvidePostgres(cfg config.Config) *postgres.Postgres {
	return postgres.NewPostgres(cfg.Postgres)
}

func ProvideMailer(cfg config.Config) mailer.Sender {
	return mailer.NewMailer(mailer.Config{
		Host:        cfg.Mailer.Host,
		Username:    cfg.Mailer.Username,
		Password:    cfg.Mailer.Password(),
		Port:        cfg.Mailer.Port,
		From:        cfg.Mailer.From(),
		TLS:         consts.MailerTLS,
		PoolConn:    consts.MailerPoolConn,
		PoolTimeout: consts.MailerPoolTimeout,
	})
}

func ProvideMailerObserver(cfg config.Config, l logger.Logger) mailer.Sender {
	return mailer.NewObserver(ProvideMailer(cfg), consts.MailerPoolConn, l)
}
