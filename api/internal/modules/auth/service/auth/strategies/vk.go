package strategies

import (
	"context"
	"io"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/vk"

	"primedivident/internal/config"
	"primedivident/internal/modules/auth/repository"
	"primedivident/internal/modules/auth/service/auth"
)

const OauthVkUrlAPI = "https://api.vk.com/method/users.get?v=5.131&album_id=wall"

type vkStrategy struct {
	oauth      *oauth2.Config
	repository repository.Repository
}

func NewVkStrategy(
	cfg config.VkOAuth2,
	repository repository.Repository,
) NetworkStrategy {
	return vkStrategy{
		oauth: &oauth2.Config{
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			RedirectURL:  cfg.RedirectUrl(),
			Scopes:       cfg.Scopes,
			Endpoint:     vk.Endpoint,
		},
		repository: repository,
	}
}

func (v vkStrategy) Callback(state string) string {
	return v.oauth.AuthCodeURL(state, oauth2.AccessTypeOnline)
}

func (v vkStrategy) Login(code string) (auth.Tokens, error) {
	token, err := v.oauth.Exchange(context.Background(), code)
	if err != nil {
		return auth.Tokens{}, err
	}

	client := v.oauth.Client(context.Background(), token)

	response, err := client.Get(OauthVkUrlAPI)
	if err != nil {
		return auth.Tokens{}, err
	}

	defer response.Body.Close()

	if _, err = io.ReadAll(response.Body); err != nil {
		return auth.Tokens{}, err
	}

	// TODO: save db refresh token

	return auth.Tokens{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}

func (v vkStrategy) Refresh(refreshToken string) (auth.Tokens, error) {
	return auth.Tokens{}, nil
}

func (v vkStrategy) Logout(refreshToken string) error {
	return nil
}
