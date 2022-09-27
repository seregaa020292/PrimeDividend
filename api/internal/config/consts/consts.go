package consts

import "time"

const (
	ServerAddr              = ":3000"
	ServerReadHeaderTimeout = 60 * time.Second

	RequestLimit = 100
	WindowLength = 1 * time.Minute

	MailerTLS         = true
	MailerPoolConn    = 4
	MailerPoolTimeout = 10 * time.Second

	TemplateCache         = true
	TemplateBaseDir       = "templates"
	TemplateMailConfirmed = "mail/views/confirmed.html"

	TimestampFormat = "02.01.2006 15:04:05"
	Timezone        = "Europe/Moscow"

	TmpLog = "./tmp/logs/server.log"

	TimeoutShutdown = 10 * time.Second

	TokenJoinTTL = 1 * time.Hour

	VkOauthRedirectUrl     = "/api/auth/vk/callback"
	OkOauthRedirectUrl     = "/api/auth/ok/callback"
	YandexOauthRedirectUrl = "/api/auth/yandex/callback"
)
