package strategies

import (
	"primedivident/internal/modules/auth/dto"
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

func (e emailStrategy) Login(email, password string, accountability entity.Accountability) (auth.Tokens, error) {
	user, err := e.repository.FindUserByEmail(email)
	if err != nil {
		return auth.Tokens{}, errorn.ErrSelect.Wrap(err)
	}

	if err := user.ErrorIsEmpty(); err != nil {
		return auth.Tokens{}, errorn.ErrNotFound.Wrap(err)
	}

	jwtPayload, err := user.JwtPayloadValidPassword(password)
	if err != nil {
		return auth.Tokens{}, errorn.ErrForbidden.Wrap(err)
	}

	genTokens, err := e.jwtTokens.GenTokens(jwtPayload)
	if err != nil {
		return auth.Tokens{}, errorn.ErrUnknown.Wrap(err)
	}

	if err := e.repository.SaveRefreshToken(dto.ModelSessionCreating(
		user.ID,
		auth.Email,
		genTokens.RefreshToken,
		accountability,
	)); err != nil {
		return auth.Tokens{}, err
	}

	if err := e.repository.RemoveExpireRefreshToken(user.ID); err != nil {
		return auth.Tokens{}, err
	}

	if err := e.repository.RemoveLastRefreshToken(user.ID); err != nil {
		return auth.Tokens{}, err
	}

	return genTokens, nil
}
