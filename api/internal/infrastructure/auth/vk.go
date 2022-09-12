package auth

import (
	"primedivident/internal/config"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/vk"
)

const oauthVkUrlAPI = "https://api.vk.com/method/users.get?v=5.131&album_id=wall"

var VkOAuth2Config = &oauth2.Config{
	ClientID:     config.GetConfig().Networks.VkOAuth2.ClientID,
	ClientSecret: config.GetConfig().Networks.VkOAuth2.ClientSecret,
	RedirectURL:  config.GetConfig().Networks.VkOAuth2.RedirectUrl(),
	Scopes:       config.GetConfig().Networks.VkOAuth2.Scopes,
	Endpoint:     vk.Endpoint,
}
