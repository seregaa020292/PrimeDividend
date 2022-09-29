package strategies

import (
	"primedivident/internal/modules/auth/entity"
	"primedivident/internal/modules/auth/service/strategy"
	"primedivident/internal/modules/auth/service/strategy/auth"
	"primedivident/internal/modules/auth/service/strategy/categorize"
	"primedivident/pkg/errorn"
)

type emailStrategy struct {
	strategy.Service
}

func NewEmailStrategy(service strategy.Service) categorize.PasswordStrategy {
	return emailStrategy{Service: service}
}

func (e emailStrategy) Login(email, password string, accountability entity.Accountability) (auth.Tokens, error) {
	user, err := e.Repository.FindUserByEmail(email)
	if err != nil {
		return auth.Tokens{}, errorn.ErrSelect.Wrap(err)
	}

	if err := user.ErrorIsEmpty(); err != nil {
		return auth.Tokens{}, errorn.ErrNotFound.Wrap(err)
	}

	if err = user.ValidPasswordActive(password); err != nil {
		return auth.Tokens{}, errorn.ErrForbidden.Wrap(err)
	}

	return e.CreateSessionTokens(auth.Email, user, accountability)
}
