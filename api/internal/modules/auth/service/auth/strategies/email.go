package strategies

import (
	"primedivident/internal/modules/auth/repository"
	"primedivident/internal/modules/auth/service/auth"
	"primedivident/pkg/errorn"
)

type emailStrategy struct {
	jwtTokens  auth.JwtTokens
	repository repository.Repository
}

func NewEmailStrategy(
	jwtTokens auth.JwtTokens,
	repository repository.Repository,
) EmailStrategy {
	return emailStrategy{
		jwtTokens:  jwtTokens,
		repository: repository,
	}
}

func (e emailStrategy) Login(email, password string) (auth.Tokens, error) {
	user, err := e.repository.FindByEmail(email)
	if err != nil {
		return auth.Tokens{}, errorn.ErrorSelect.Wrap(err)
	}

	if !user.Status.IsActive() {
		return auth.Tokens{}, errorn.ErrorAccess
	}

	if err := user.ComparePasswordHash(password); err != nil {
		return auth.Tokens{}, errorn.ErrorPasswordIncorrect.Wrap(err)
	}

	genTokens, err := e.jwtTokens.GenTokens(user.JwtPayload())
	if err != nil {
		return auth.Tokens{}, errorn.ErrorUnknown.Wrap(err)
	}

	// TODO: save db refresh token

	return genTokens, nil
}

func (e emailStrategy) Validate(token string) error {
	_, err := e.jwtTokens.ValidateAccessToken(token)

	return err
}

func (e emailStrategy) Refresh(refreshToken string) (auth.Tokens, error) {
	return auth.Tokens{}, nil
}

func (e emailStrategy) Logout(refreshToken string) error {
	return nil
}
