package strategies

import (
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

func (s emailStrategy) Login(email, password string, accountability auth.Accountability) (auth.Tokens, error) {
	user, err := s.Repository.FindUserByEmail(email)
	if err != nil {
		return auth.Tokens{}, errorn.ErrSelect.Wrap(err)
	}

	if err := user.ErrorIsEmpty(); err != nil {
		return auth.Tokens{}, errorn.ErrNotFound.Wrap(err)
	}

	if err = user.ValidPasswordActive(password); err != nil {
		return auth.Tokens{}, errorn.ErrForbidden.Wrap(err)
	}

	return s.CreateSessionTokens(user, accountability)
}
