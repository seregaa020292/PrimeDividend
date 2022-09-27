package strategy

import (
	"primedivident/internal/modules/auth/service/strategy/auth"
	"primedivident/internal/modules/auth/service/strategy/categorize"
	"primedivident/internal/modules/auth/service/strategy/repository"
)

type Strategy interface {
	Network() categorize.NetworkStrategies
	Password() categorize.PasswordStrategies
	Verify(accessToken string) error
	Logout(refreshToken string) error
	Refresh(refreshToken string) (auth.Tokens, error)
}

type strategy struct {
	categorize categorize.Categorize
	jwtTokens  auth.JwtTokens
	repository repository.Repository
}

func NewStrategy(jwtTokens auth.JwtTokens, repository repository.Repository) Strategy {
	return &strategy{
		categorize: categorize.NewCategorize(),
		jwtTokens:  jwtTokens,
		repository: repository,
	}
}

func (s strategy) Network() categorize.NetworkStrategies {
	return s.categorize.Networks
}

func (s strategy) Password() categorize.PasswordStrategies {
	return s.categorize.Passwords
}

func (s strategy) Verify(accessToken string) error {
	_, err := s.jwtTokens.ValidateAccessToken(accessToken)

	return err
}

func (s strategy) Logout(refreshToken string) error {
	return nil
}

func (s strategy) Refresh(refreshToken string) (auth.Tokens, error) {
	return auth.Tokens{}, nil
}
