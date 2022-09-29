package strategies

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/vk"

	"primedivident/internal/config"
	"primedivident/internal/modules/auth/dto"
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

func (v vkStrategy) Login(code string, accountability entity.Accountability) (auth.Tokens, error) {
	network, err := v.authorization(code)
	if err != nil {
		return auth.Tokens{}, errorn.ErrUnauthorized.Wrap(err)
	}

	user, err := v.repository.FindUserByEmail(network.Email)
	if err != nil {
		return auth.Tokens{}, errorn.ErrSelect.Wrap(err)
	}

	if user.IsEmpty() {
		user = entity.NewUserNetwork(network.Email, network.Name)

		if err := v.repository.CreateUser(dto.ModelUserByEntity(user)); err != nil {
			return auth.Tokens{}, errorn.ErrInsert.Wrap(err)
		}
	}

	if err := user.ErrorIsActiveStatus(); err != nil {
		return auth.Tokens{}, errorn.ErrNotFound.Wrap(err)
	}

	userNetwork, err := v.repository.FindNetworkByID(network.Identity, auth.Vk)
	if err != nil {
		return auth.Tokens{}, errorn.ErrUnknown.Wrap(err)
	}

	if userNetwork.IsEmpty() {
		if err := v.repository.AttachNetwork(dto.ModelUserNetworksCreating(network, user.ID, auth.Vk)); err != nil {
			return auth.Tokens{}, errorn.ErrInsert.Wrap(err)
		}
	}

	genTokens, err := v.jwtTokens.GenTokens(user.JwtPayload())
	if err != nil {
		return auth.Tokens{}, errorn.ErrUnknown.Wrap(err)
	}

	if err := v.repository.SaveRefreshToken(dto.ModelSessionCreating(
		user.ID,
		auth.Vk,
		genTokens.RefreshToken,
		accountability,
	)); err != nil {
		return auth.Tokens{}, err
	}

	if err := v.repository.RemoveExpireRefreshToken(user.ID); err != nil {
		return auth.Tokens{}, err
	}

	if err := v.repository.RemoveLastRefreshToken(user.ID); err != nil {
		return auth.Tokens{}, err
	}

	return genTokens, nil
}

func (v vkStrategy) authorization(code string) (entity.Network, error) {
	token, err := v.oauth.Exchange(context.Background(), code)
	if err != nil {
		return entity.Network{}, errorn.ErrUnknown.Wrap(err)
	}

	client := v.oauth.Client(context.Background(), token)

	response, err := client.Get(OauthVkUrlAPI)
	if err != nil {
		return entity.Network{}, errorn.ErrUnknown.Wrap(err)
	}
	defer response.Body.Close()

	var body vkBody
	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		return entity.Network{}, errorn.ErrUnknown.Wrap(err)
	}
	bodyResponse := body.Response[0]

	return entity.Network{
		Identity: strconv.Itoa(bodyResponse.ID),
		Email:    token.Extra("email").(string),
		Name:     fmt.Sprintf("%s %s", bodyResponse.LastName, bodyResponse.FirstName),
	}, nil
}
