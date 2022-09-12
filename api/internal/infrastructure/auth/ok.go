package auth

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/odnoklassniki"

	"primedivident/internal/config"
)

func OkOAuth2Config(cfg config.OkOAuth2) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		Endpoint:     odnoklassniki.Endpoint,
		RedirectURL:  cfg.RedirectUrl(),
		Scopes:       cfg.Scopes,
	}
}
