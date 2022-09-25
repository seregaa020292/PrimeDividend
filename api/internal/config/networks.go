package config

import (
	"net/url"

	"primedivident/internal/config/consts"
)

type Networks struct {
	VkOAuth2     VkOAuth2
	YandexOAuth2 YandexOAuth2
	OkOAuth2     OkOAuth2
}

type VkOAuth2 struct {
	ClientID     string   `env:"VK_OAUTH_CLIENT_ID" env-required:"true"`
	ClientSecret string   `env:"VK_OAUTH_CLIENT_SECRET" env-required:"true"`
	Scopes       []string `env:"VK_OAUTH_SCOPES"`

	SiteOrigin string `env:"SITE_ORIGIN" env-required:"true"`
}

type YandexOAuth2 struct {
	ClientID     string   `env:"YANDEX_OAUTH_CLIENT_ID" env-required:"true"`
	ClientSecret string   `env:"YANDEX_OAUTH_CLIENT_SECRET" env-required:"true"`
	Scopes       []string `env:"YANDEX_OAUTH_SCOPES"`

	SiteOrigin string `env:"SITE_ORIGIN" env-required:"true"`
}

type OkOAuth2 struct {
	ClientID     string   `env:"OK_OAUTH_CLIENT_ID" env-required:"true"`
	ClientSecret string   `env:"OK_OAUTH_CLIENT_SECRET" env-required:"true"`
	Scopes       []string `env:"OK_OAUTH_SCOPES"`

	SiteOrigin string `env:"SITE_ORIGIN" env-required:"true"`
}

func (a VkOAuth2) RedirectUrl() string {
	path, _ := url.JoinPath(a.SiteOrigin, consts.VkOauthRedirectUrl)
	return path
}

func (a YandexOAuth2) RedirectUrl() string {
	path, _ := url.JoinPath(a.SiteOrigin, consts.YandexOauthRedirectUrl)
	return path
}

func (a OkOAuth2) RedirectUrl() string {
	path, _ := url.JoinPath(a.SiteOrigin, consts.OkOauthRedirectUrl)
	return path
}
