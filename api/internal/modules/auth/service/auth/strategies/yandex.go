package strategies

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/yandex"

	"primedivident/internal/config"
	"primedivident/internal/modules/auth/repository"
	"primedivident/internal/modules/auth/service/auth"
)

type yandexStrategy struct {
	oauth      *oauth2.Config
	repository repository.Repository
}

func NewYandexStrategy(
	cfg config.YandexOAuth2,
	repository repository.Repository,
) NetworkStrategy {
	return yandexStrategy{
		oauth: &oauth2.Config{
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			Endpoint:     yandex.Endpoint,
			RedirectURL:  cfg.RedirectUrl(),
			Scopes:       cfg.Scopes,
		},
		repository: repository,
	}
}

func (y yandexStrategy) Callback(state string) string {
	return y.oauth.AuthCodeURL(state, oauth2.AccessTypeOnline)
}

func (y yandexStrategy) Login(code string) (auth.Tokens, error) {
	//TODO implement me
	panic("implement me")
}

func (y yandexStrategy) Validate(token string) error {
	return nil
}

func (y yandexStrategy) Refresh(refreshToken string) (auth.Tokens, error) {
	return auth.Tokens{}, nil
}

func (y yandexStrategy) Logout(refreshToken string) error {
	return nil
}
