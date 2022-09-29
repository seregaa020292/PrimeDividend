package strategies

import (
	"fmt"
	"strconv"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/vk"

	"primedivident/internal/config"
	"primedivident/internal/modules/auth/entity"
	"primedivident/internal/modules/auth/service/strategy"
	"primedivident/internal/modules/auth/service/strategy/auth"
	"primedivident/internal/modules/auth/service/strategy/categorize"
	"primedivident/pkg/errorn"
)

const OauthVkUrlAPI = "https://api.vk.com/method/users.get?v=5.131&album_id=wall"

type vkStrategy struct {
	oauth *oauth2.Config
	strategy.Service
}

func NewVkStrategy(cfg config.VkOAuth2, service strategy.Service) categorize.NetworkStrategy {
	return vkStrategy{
		oauth: &oauth2.Config{
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			RedirectURL:  cfg.RedirectUrl(),
			Scopes:       cfg.Scopes,
			Endpoint:     vk.Endpoint,
		},
		Service: service,
	}
}

func (v vkStrategy) Callback(state string) string {
	return v.oauth.AuthCodeURL(state, oauth2.AccessTypeOnline)
}

func (v vkStrategy) Login(code string, accountability entity.Accountability) (auth.Tokens, error) {
	var response responseVK

	token, err := v.ClientNetwork(&response, v.oauth, code, OauthVkUrlAPI)
	if err != nil {
		return auth.Tokens{}, err
	}

	network := entity.Network{
		Identity: strconv.Itoa(response.Response[0].ID),
		Email:    token.Extra("email").(string),
		Name:     fmt.Sprintf("%s %s", response.Response[0].LastName, response.Response[0].FirstName),
	}

	user, err := v.UserAttachNetwork(network, auth.Vk)
	if err != nil {
		return auth.Tokens{}, errorn.ErrUnauthorized.Wrap(err)
	}

	return v.CreateSessionTokens(auth.Vk, user, accountability)
}
