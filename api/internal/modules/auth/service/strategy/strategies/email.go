package strategies

import (
	"primedivident/internal/models/app/public/model"
	"primedivident/internal/modules/auth/entity"
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

func (e emailStrategy) Login(email, password string, session entity.FingerprintSession) (auth.Tokens, error) {
	user, err := e.repository.FindUserByEmail(email)
	if err != nil {
		return auth.Tokens{}, errorn.ErrSelect.Wrap(err)
	}

	jwtPayload, err := user.JwtPayloadValidPassword(password)
	if err != nil {
		return auth.Tokens{}, errorn.ErrForbidden.Wrap(err)
	}

	genTokens, err := e.jwtTokens.GenTokens(jwtPayload)
	if err != nil {
		return auth.Tokens{}, errorn.ErrUnknown.Wrap(err)
	}

	if err := e.repository.SaveRefreshToken(model.Sessions{
		Token:     &genTokens.RefreshToken.Value,
		ExpiresAt: genTokens.RefreshToken.ExpiresAt,
		UserID:    user.ID,
		Strategy:  auth.Email.String(),
		IP:        session.IP,
		UserAgent: session.UserAgent,
	}); err != nil {
		return auth.Tokens{}, err
	}

	return genTokens, nil
}
