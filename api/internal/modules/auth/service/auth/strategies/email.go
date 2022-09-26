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
) auth.PasswordStrategy {
	return emailStrategy{
		jwtTokens:  jwtTokens,
		repository: repository,
	}
}

func (e emailStrategy) Login(email, password string) (auth.Tokens, error) {
	user, err := e.repository.FindByEmail(email)
	if err != nil {
		return auth.Tokens{}, errorn.ErrSelect.Wrap(err)
	}

	if !user.Status.IsActive() {
		return auth.Tokens{}, errorn.ErrForbidden
	}

	if err := user.ComparePasswordHash(password); err != nil {
		return auth.Tokens{}, errorn.ErrPasswordIncorrect.Wrap(err)
	}

	genTokens, err := e.jwtTokens.GenTokens(user.JwtPayload())
	if err != nil {
		return auth.Tokens{}, errorn.ErrUnknown.Wrap(err)
	}

	// TODO: save db refresh token

	return genTokens, nil
}
