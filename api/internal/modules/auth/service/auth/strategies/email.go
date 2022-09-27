package strategies

import (
	"primedivident/internal/modules/auth/service/auth"
	"primedivident/pkg/errorn"
)

type emailStrategy struct {
	jwtTokens  auth.JwtTokens
	repository auth.TokenRepository
}

func NewEmailStrategy(
	jwtTokens auth.JwtTokens,
	repository auth.TokenRepository,
) auth.PasswordStrategy {
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
