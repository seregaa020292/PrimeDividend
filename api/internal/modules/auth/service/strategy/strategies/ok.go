package strategies

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/odnoklassniki"

	"primedivident/internal/config"
	"primedivident/internal/modules/auth/entity"
	"primedivident/internal/modules/auth/service/strategy"
	"primedivident/internal/modules/auth/service/strategy/auth"
	"primedivident/internal/modules/auth/service/strategy/categorize"
)

type okStrategy struct {
	oauth *oauth2.Config
	strategy.Service
}

func NewOkStrategy(cfg config.OkOAuth2, service strategy.Service) categorize.NetworkStrategy {
	return okStrategy{
		oauth: &oauth2.Config{
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			RedirectURL:  cfg.RedirectUrl(),
			Scopes:       cfg.Scopes,
			Endpoint:     odnoklassniki.Endpoint,
		},
		Service: service,
	}
}

func (o okStrategy) Callback(state string) string {
	return o.oauth.AuthCodeURL(state, oauth2.AccessTypeOnline)
}

func (o okStrategy) Login(code string, accountability entity.Accountability) (auth.Tokens, error) {
	panic("implement me")
}
