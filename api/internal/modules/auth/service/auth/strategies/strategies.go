package strategies

import (
	"errors"

	"primedivident/internal/config"
	"primedivident/internal/modules/auth/repository"
	"primedivident/internal/modules/auth/service/auth"
)

const (
	Email  Key = "email"
	Vk     Key = "vk"
	Ok     Key = "ok"
	Yandex Key = "yandex"
)

type (
	Strategies interface {
		Email() EmailStrategy
		Networks() NetworkStrategies
		Verify(accessToken string) error
		Refresh(refreshToken string) (auth.Tokens, error)
		Logout(refreshToken string) error
	}
	NetworkStrategy interface {
		Callback(state string) string
		Login(code string) (auth.Tokens, error)
	}
	EmailStrategy interface {
		Login(email, password string) (auth.Tokens, error)
	}
)

type (
	Key               string
	NetworkStrategies map[Key]NetworkStrategy
	strategies        struct {
		email     EmailStrategy
		networks  NetworkStrategies
		jwtTokens auth.JwtTokens
	}
)

func NewStrategies(
	cfg config.Networks,
	jwtTokens auth.JwtTokens,
	repository repository.Repository,
) Strategies {
	return strategies{
		email: NewEmailStrategy(jwtTokens, repository),
		networks: NetworkStrategies{
			Vk:     NewVkStrategy(cfg.VkOAuth2, jwtTokens, repository),
			Ok:     NewOkStrategy(cfg.OkOAuth2, jwtTokens, repository),
			Yandex: NewYandexStrategy(cfg.YandexOAuth2, jwtTokens, repository),
		},
		jwtTokens: jwtTokens,
	}
}

func (s strategies) Email() EmailStrategy {
	return s.email
}

func (s strategies) Networks() NetworkStrategies {
	return s.networks
}

func (s strategies) Verify(accessToken string) error {
	_, err := s.jwtTokens.ValidateAccessToken(accessToken)

	return err
}

func (s strategies) Refresh(refreshToken string) (auth.Tokens, error) {
	return auth.Tokens{}, nil
}

func (s strategies) Logout(refreshToken string) error {
	return nil
}

func (n NetworkStrategies) Strategies() []NetworkStrategy {
	strategies := make([]NetworkStrategy, 0, len(n))

	for _, strategy := range n {
		strategies = append(strategies, strategy)
	}

	return strategies
}

func (n NetworkStrategies) SetStrategy(name Key, strategy NetworkStrategy) error {
	if _, ok := n[name]; ok {
		return errors.New("strategy already exist")
	}

	n[name] = strategy

	return nil
}

func (n NetworkStrategies) GetStrategy(name Key) NetworkStrategy {
	if strategy, ok := n[name]; ok {
		return strategy
	}

	return nil
}
