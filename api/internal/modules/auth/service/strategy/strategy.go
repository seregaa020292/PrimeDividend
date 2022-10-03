package strategy

import (
	"primedivident/internal/modules/auth/service/strategy/auth"
	"primedivident/internal/modules/auth/service/strategy/categorize"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
)

type Strategy interface {
	Network() categorize.NetworkStrategies
	Password() categorize.PasswordStrategies
	VerifyAccess(accessToken string) error
	VerifyRefresh(refreshToken string) error
	Logout(refreshToken string) error
	Refresh(refreshToken string, accountability auth.Accountability) (auth.Tokens, error)
}

type strategy struct {
	categorize categorize.Categorize
	service    Service
}

func NewStrategy(service Service) Strategy {
	return &strategy{
		categorize: categorize.NewCategorize(),
		service:    service,
	}
}

func (s strategy) Network() categorize.NetworkStrategies {
	return s.categorize.Networks
}

func (s strategy) Password() categorize.PasswordStrategies {
	return s.categorize.Passwords
}

func (s strategy) VerifyAccess(accessToken string) error {
	_, err := s.service.JwtTokens.ValidateAccessToken(accessToken)

	return errs.Forbidden.Wrap(err, errmsg.AccessDenied)
}

func (s strategy) VerifyRefresh(refreshToken string) error {
	if _, err := s.service.JwtTokens.ValidateRefreshToken(refreshToken); err != nil {
		if err := s.service.Repository.RemoveRefreshToken(refreshToken); err != nil {
			return errs.BadRequest.Wrap(err, errmsg.FailedDeleteData)
		}
		return errs.BadRequest.Wrap(err, errmsg.ConfirmWhileMatching)
	}

	return nil
}

func (s strategy) Logout(refreshToken string) error {
	if err := s.service.JwtTokens.CorrectRefreshToken(refreshToken); err != nil {
		return errs.BadRequest.Wrap(err, errmsg.CheckingWhileOccurred)
	}

	if err := s.service.Repository.RemoveRefreshToken(refreshToken); err != nil {
		return errs.BadRequest.Wrap(err, errmsg.FailedDeleteData)
	}

	return nil
}

func (s strategy) Refresh(refreshToken string, accountability auth.Accountability) (auth.Tokens, error) {
	tokens, err := s.service.UpdateSessionTokens(refreshToken, accountability)
	if err != nil {
		return auth.Tokens{}, errs.BadRequest.Wrap(err, errmsg.FailedUpdateData)
	}

	return tokens, nil
}
