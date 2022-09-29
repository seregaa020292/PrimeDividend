package strategy

import (
	"primedivident/internal/modules/auth/entity"
	"primedivident/internal/modules/auth/service/strategy/auth"
	"primedivident/internal/modules/auth/service/strategy/categorize"
)

type Strategy interface {
	Network() categorize.NetworkStrategies
	Password() categorize.PasswordStrategies
	Verify(accessToken string) error
	Logout(refreshToken string) error
	Refresh(refreshToken string, accountability entity.Accountability) (auth.Tokens, error)
}

type strategy struct {
	categorize categorize.Categorize
	service    Service
}

func NewStrategy(service Service) Strategy {
	return &strategy{
		categorize: categorize.NewCategorize(),
		service:    service,
	}
}

func (s strategy) Network() categorize.NetworkStrategies {
	return s.categorize.Networks
}

func (s strategy) Password() categorize.PasswordStrategies {
	return s.categorize.Passwords
}

func (s strategy) Verify(accessToken string) error {
	_, err := s.service.JwtTokens.ValidateAccessToken(accessToken)

	return err
}

func (s strategy) Logout(refreshToken string) error {
	if _, err := s.service.JwtTokens.ValidateRefreshToken(refreshToken); err != nil {
		return err
	}

	return s.service.Repository.RemoveRefreshToken(refreshToken)
}

func (s strategy) Refresh(refreshToken string, accountability entity.Accountability) (auth.Tokens, error) {
	panic("implement me")
}
