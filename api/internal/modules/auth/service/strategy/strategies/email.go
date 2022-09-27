package strategies

import (
	"primedivident/internal/modules/auth/service/strategy/auth"
	"primedivident/internal/modules/auth/service/strategy/categorize"
	"primedivident/internal/modules/auth/service/strategy/repository"
	"primedivident/pkg/errorn"
)

type emailStrategy struct {
	jwtTokens  auth.JwtTokens
	repository repository.Repository
}

func NewEmailStrategy(jwtTokens auth.JwtTokens, repository repository.Repository) categorize.PasswordStrategy {
	return emailStrategy{
		jwtTokens:  jwtTokens,
		repository: repository,
	}
}

func (e emailStrategy) Login(email, password string) (auth.Tokens, error) {
	user, err := e.repository.FindUserByEmail(email)
	if err != nil {
		return auth.Tokens{}, errorn.ErrSelect.Wrap(err)
	}

	jwtUser, err := user.JwtPayloadValidPassword(password)
	if err != nil {
		return auth.Tokens{}, errorn.ErrForbidden.Wrap(err)
	}

	genTokens, err := e.jwtTokens.GenTokens(jwtUser)
	if err != nil {
		return auth.Tokens{}, errorn.ErrUnknown.Wrap(err)
	}

	e.repository.SaveRefreshToken(user.ID, genTokens.RefreshToken)

	return genTokens, nil
}
