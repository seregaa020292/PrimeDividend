package strategies

import (
	"primedivident/internal/modules/auth/service/strategy"
	"primedivident/internal/modules/auth/service/strategy/auth"
	"primedivident/internal/modules/auth/service/strategy/categorize"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
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
		return auth.Tokens{}, errs.BadRequest.Wrap(err, errmsg.FailedGetData)
	}

	if err := user.ErrorIsEmpty(); err != nil {
		return auth.Tokens{}, errs.BadRequest.Wrap(err, errmsg.FailedGetData)
	}

	if err = user.ValidPasswordActive(password); err != nil {
		return auth.Tokens{}, errs.BadRequest.Wrap(err, errmsg.CheckingWhileOccurred)
	}

	tokens, err := s.CreateSessionTokens(user, accountability)
	if err != nil {
		return auth.Tokens{}, errs.BadRequest.Wrap(err, errmsg.FailedAddData)
	}

	return tokens, nil
}
