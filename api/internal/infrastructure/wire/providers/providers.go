package providers

import (
	"primedivident/internal/config"
	"primedivident/internal/config/consts"
	"primedivident/internal/modules/auth/service/strategy"
	"primedivident/internal/modules/auth/service/strategy/auth"
	"primedivident/internal/modules/auth/service/strategy/strategies"
	"primedivident/internal/modules/market/service/quotes"
	"primedivident/internal/modules/market/service/quotes/providers"
	"primedivident/pkg/db/postgres"
	"primedivident/pkg/db/redis"
	"primedivident/pkg/graceful"
	"primedivident/pkg/logger"
	"primedivident/pkg/mailer"
	"primedivident/pkg/tpl"
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

func ProvideRedis(cfg config.Config) *redis.Redis {
	return redis.NewRedis(cfg.Redis)
}

func ProvideMailerObserver(cfg config.Config, log logger.Logger) mailer.Sender {
	return mailer.NewObserver(mailer.NewMailer(mailer.Config{
		Host:        cfg.Mailer.Host,
		Username:    cfg.Mailer.Username,
		Password:    cfg.Mailer.Password(),
		Port:        cfg.Mailer.Port,
		From:        cfg.Mailer.From(),
		TLS:         consts.MailerTLS,
		PoolConn:    consts.MailerPoolConn,
		PoolTimeout: consts.MailerPoolTimeout,
	}), consts.MailerPoolConn, log)
}

func ProvideJwtTokens(cfg config.Config) auth.JwtTokens {
	return auth.NewJwtTokens(cfg.App.Name, cfg.Jwt)
}

func ProvideTemplate(cfg config.Config) tpl.Templater {
	return tpl.NewTemplate(consts.TemplateBaseDir, consts.TemplateCache, map[string]any{
		"siteOrigin": cfg.App.SiteOrigin,
		"appName":    cfg.App.Name,
	})
}

func ProvideStrategy(cfg config.Config, service strategy.Service) strategy.Strategy {
	strategics := strategy.NewStrategy(service)

	strategics.Password().Set(auth.Email, strategies.NewEmailStrategy(service))
	strategics.Network().Set(auth.Vk, strategies.NewVkStrategy(cfg.Networks.VkOAuth2, service))
	strategics.Network().Set(auth.Ok, strategies.NewOkStrategy(cfg.Networks.OkOAuth2, service))
	strategics.Network().Set(auth.Yandex, strategies.NewYandexStrategy(cfg.Networks.YandexOAuth2, service))

	return strategics
}

func ProvideTinkoff(config config.Config) providers.Tinkoff {
	return providers.NewTinkoff(config.Tinkoff)
}

func ProvideShutdown(
	postgres *postgres.Postgres,
	redis *redis.Redis,
	mailer mailer.Sender,
	hubQuotes *quotes.HubQuotes,
) graceful.ShutdownApp {
	return graceful.ShutdownApp{
		mailer.Close,
		hubQuotes.Close,
		redis.Close,
		postgres.Close,
	}
}
