package strategies

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/odnoklassniki"

	"primedivident/internal/config"
	"primedivident/internal/modules/auth/repository"
	"primedivident/internal/modules/auth/service/auth"
)

type okStrategy struct {
	oauth      *oauth2.Config
	repository repository.Repository
}

func NewOkStrategy(
	cfg config.OkOAuth2,
	repository repository.Repository,
) NetworkStrategy {
	return okStrategy{
		oauth: &oauth2.Config{
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			RedirectURL:  cfg.RedirectUrl(),
			Scopes:       cfg.Scopes,
			Endpoint:     odnoklassniki.Endpoint,
		},
		repository: repository,
	}
}

func (o okStrategy) Callback(state string) string {
	return o.oauth.AuthCodeURL(state, oauth2.AccessTypeOnline)
}

func (o okStrategy) Login(code string) (auth.Tokens, error) {
	//TODO implement me
	panic("implement me")
}

func (o okStrategy) Refresh(refreshToken string) (auth.Tokens, error) {
	return auth.Tokens{}, nil
}

func (o okStrategy) Logout(refreshToken string) error {
	return nil
}
