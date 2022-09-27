package strategies

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/odnoklassniki"

	"primedivident/internal/config"
	"primedivident/internal/modules/auth/entity"
	"primedivident/internal/modules/auth/service/strategy/categorize"
	"primedivident/internal/modules/auth/service/strategy/repository"
)

type okStrategy struct {
	oauth      *oauth2.Config
	jwtTokens  entity.JwtTokens
	repository repository.Repository
}

func NewOkStrategy(cfg config.OkOAuth2, jwtTokens entity.JwtTokens, repository repository.Repository) categorize.NetworkStrategy {
	return okStrategy{
		oauth: &oauth2.Config{
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			RedirectURL:  cfg.RedirectUrl(),
			Scopes:       cfg.Scopes,
			Endpoint:     odnoklassniki.Endpoint,
		},
		jwtTokens:  jwtTokens,
		repository: repository,
	}
}

func (o okStrategy) Callback(state string) string {
	return o.oauth.AuthCodeURL(state, oauth2.AccessTypeOnline)
}

func (o okStrategy) Login(code string) (entity.Tokens, error) {
	panic("implement me")
}
