package auth

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/yandex"

	"primedivident/internal/config"
)

func YandexOAuth2Config(cfg config.YandexOAuth2) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		Endpoint:     yandex.Endpoint,
		RedirectURL:  cfg.RedirectUrl(),
		Scopes:       cfg.Scopes,
	}
}
