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
	NetworkStrategy interface {
		Callback(state string) string
		Login(code string) (auth.Tokens, error)
		Refresh
		Logout
	}
	EmailStrategy interface {
		Login(email, password string) (auth.Tokens, error)
		Refresh
		Logout
	}
	Refresh interface {
		Refresh(refreshToken string) (auth.Tokens, error)
	}
	Logout interface {
		Logout(refreshToken string) error
	}
)

type (
	Key               string
	NetworkStrategies map[Key]NetworkStrategy
	Strategies        struct {
		Email    EmailStrategy
		Networks NetworkStrategies
	}
)

func NewStrategies(
	cfg config.Networks,
	jwtTokens auth.JwtTokens,
	repository repository.Repository,
) Strategies {
	return Strategies{
		Email: NewEmailStrategy(jwtTokens, repository),
		Networks: NetworkStrategies{
			Vk:     NewVkStrategy(cfg.VkOAuth2, repository),
			Ok:     NewOkStrategy(cfg.OkOAuth2, repository),
			Yandex: NewYandexStrategy(cfg.YandexOAuth2, repository),
		},
	}
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
