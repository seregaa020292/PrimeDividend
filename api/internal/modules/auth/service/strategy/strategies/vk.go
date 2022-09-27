package strategies

import (
	"context"
	"encoding/json"
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/vk"

	"primedivident/internal/config"
	"primedivident/internal/modules/auth/entity"
	"primedivident/internal/modules/auth/service/strategy/auth"
	"primedivident/internal/modules/auth/service/strategy/categorize"
	"primedivident/internal/modules/auth/service/strategy/repository"
	"primedivident/pkg/errorn"
)

const OauthVkUrlAPI = "https://api.vk.com/method/users.get?v=5.131&album_id=wall"

type vkStrategy struct {
	oauth      *oauth2.Config
	jwtTokens  auth.JwtTokens
	repository repository.Repository
}

func NewVkStrategy(
	cfg config.VkOAuth2,
	jwtTokens auth.JwtTokens,
	repository repository.Repository,
) categorize.NetworkStrategy {
	return vkStrategy{
		oauth: &oauth2.Config{
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			RedirectURL:  cfg.RedirectUrl(),
			Scopes:       cfg.Scopes,
			Endpoint:     vk.Endpoint,
		},
		jwtTokens:  jwtTokens,
		repository: repository,
	}
}

func (v vkStrategy) Callback(state string) string {
	return v.oauth.AuthCodeURL(state, oauth2.AccessTypeOnline)
}

func (v vkStrategy) Login(code string) (auth.Tokens, error) {
	token, err := v.oauth.Exchange(context.Background(), code)
	if err != nil {
		return auth.Tokens{}, errorn.ErrUnknown.Wrap(err)
	}

	client := v.oauth.Client(context.Background(), token)

	response, err := client.Get(OauthVkUrlAPI)
	if err != nil {
		return auth.Tokens{}, errorn.ErrUnknown.Wrap(err)
	}

	defer response.Body.Close()

	var body vkBody
	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		return auth.Tokens{}, errorn.ErrUnknown.Wrap(err)
	}

	jwtUser := entity.JwtUser{
		Email: token.Extra("email").(string),
		Name:  fmt.Sprintf("%s %s", body.Response[0].LastName, body.Response[0].FirstName),
	}

	genTokens, err := v.jwtTokens.GenTokens(jwtUser)
	if err != nil {
		return auth.Tokens{}, errorn.ErrUnknown.Wrap(err)
	}

	v.repository.AttachNetwork(jwtUser, auth.Vk)

	return genTokens, nil
}
