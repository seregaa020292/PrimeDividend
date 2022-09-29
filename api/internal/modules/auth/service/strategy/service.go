package strategy

import (
	"context"
	"encoding/json"

	"golang.org/x/oauth2"

	"primedivident/internal/modules/auth/dto"
	"primedivident/internal/modules/auth/entity"
	"primedivident/internal/modules/auth/service/strategy/auth"
	"primedivident/internal/modules/auth/service/strategy/repository"
)

type Service struct {
	JwtTokens  auth.JwtTokens
	Repository repository.Repository
}

func NewService(jwtTokens auth.JwtTokens, repository repository.Repository) Service {
	return Service{
		JwtTokens:  jwtTokens,
		Repository: repository,
	}
}

func (s Service) CreateSessionTokens(
	strategy auth.Name,
	user entity.User,
	accountability entity.Accountability,
) (auth.Tokens, error) {
	genTokens, err := s.JwtTokens.GenTokens(user.JwtPayload())
	if err != nil {
		return auth.Tokens{}, err
	}

	if err := s.Repository.SaveRefreshToken(dto.ModelSessionCreating(
		user.ID,
		strategy,
		genTokens.RefreshToken,
		accountability,
	)); err != nil {
		return auth.Tokens{}, err
	}

	if err := s.Repository.RemoveExpireRefreshToken(user.ID); err != nil {
		return auth.Tokens{}, err
	}

	if err := s.Repository.RemoveLastRefreshToken(user.ID); err != nil {
		return auth.Tokens{}, err
	}

	return genTokens, nil
}

func (s Service) ClientNetwork(
	body any,
	oauth *oauth2.Config,
	code string,
	urlApi func(*oauth2.Token) string,
) (*oauth2.Token, error) {
	token, err := oauth.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	client := oauth.Client(context.Background(), token)

	response, err := client.Get(urlApi(token))
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(body); err != nil {
		return nil, err
	}

	return token, err
}

func (s Service) UserAttachNetwork(network entity.Network, strategy auth.Name) (entity.User, error) {
	user, err := s.Repository.FindUserByEmail(network.Email)
	if err != nil {
		return entity.User{}, err
	}

	if user.IsEmpty() {
		user = entity.NewUserNetwork(network.Email, network.Name)

		if err := s.Repository.CreateUser(dto.ModelUserByEntity(user)); err != nil {
			return entity.User{}, err
		}
	}

	if err := user.ErrorIsActiveStatus(); err != nil {
		return entity.User{}, err
	}

	userNetwork, err := s.Repository.FindNetworkByID(network.Identity, strategy)
	if err != nil {
		return entity.User{}, err
	}

	if userNetwork.IsEmpty() {
		if err := s.Repository.AttachNetwork(dto.ModelUserNetworksCreating(network, user.ID, strategy)); err != nil {
			return entity.User{}, err
		}
	}

	return user, nil
}
