package strategies

import (
	"fmt"
	"strconv"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/vk"

	"primedividend/api/internal/config"
	"primedividend/api/internal/modules/auth/entity"
	"primedividend/api/internal/modules/auth/service/strategy"
	"primedividend/api/internal/modules/auth/service/strategy/auth"
	"primedividend/api/internal/modules/auth/service/strategy/categorize"
	"primedividend/api/pkg/errs"
	"primedividend/api/pkg/errs/errmsg"
)

const oauthUrlVK = "https://api.vk.com/method/users.get?v=5.131"

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

func (s vkStrategy) Callback(state string) string {
	return s.oauth.AuthCodeURL(state, oauth2.AccessTypeOnline)
}

func (s vkStrategy) Login(code string, accountability auth.Accountability) (auth.Tokens, error) {
	var response responseVK
	var oauthToken *oauth2.Token

	if err := s.ClientNetwork(&response, code, s.oauth, func(token *oauth2.Token) string {
		oauthToken = token
		return oauthUrlVK
	}); err != nil {
		return auth.Tokens{}, errs.BadRequest.Wrap(err, errmsg.EncounteredRequestExternal)
	}

	network := entity.Network{
		ClientID: strconv.Itoa(response.Response[0].ID),
		Email:    oauthToken.Extra("email").(string),
		Name:     fmt.Sprintf("%s %s", response.Response[0].LastName, response.Response[0].FirstName),
	}

	user, err := s.UserAttachNetwork(network, auth.Vk)
	if err != nil {
		return auth.Tokens{}, errs.BadRequest.Wrap(err, errmsg.FailedUpdateData)
	}

	tokens, err := s.CreateSessionTokens(user, accountability)
	if err != nil {
		return auth.Tokens{}, errs.BadRequest.Wrap(err, errmsg.FailedAddData)
	}

	return tokens, nil
}
