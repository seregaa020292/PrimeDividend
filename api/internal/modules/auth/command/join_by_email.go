package command

import (
	"primedivident/internal/decorator"
	"primedivident/internal/models/app/public/model"
	"primedivident/internal/modules/auth/entity"
	"primedivident/internal/modules/auth/repository"
	"primedivident/internal/modules/auth/service/email"
	"primedivident/pkg/errorn"
	"primedivident/pkg/utils/gog"
)

type (
	Credential struct {
		Email    string
		Password string
	}
	JoinByEmail decorator.CommandHandler[Credential]
)

type joinByEmail struct {
	repository repository.Repository
	email      email.JoinConfirmUser
}

func NewJoinByEmail(
	repository repository.Repository,
	email email.JoinConfirmUser,
) JoinByEmail {
	return joinByEmail{
		repository: repository,
		email:      email,
	}
}

func (c joinByEmail) Exec(cmd Credential) error {
	if user, err := c.repository.FindByEmail(cmd.Email); err != nil {
		return errorn.ErrSelect.Wrap(err)
	} else {
		if user != (entity.User{}) {
			return c.existedUser(user)
		}
	}

	user, err := entity.NewUser(cmd.Email, cmd.Password)
	if err != nil {
		return errorn.ErrUnknown.Wrap(err)
	}

	if err := c.repository.Add(model.Users{
		Email:            user.Email,
		Password:         user.PassHash,
		Role:             user.Role.String(),
		Status:           user.Status.String(),
		TokenJoinValue:   gog.Ptr(user.Token.Value),
		TokenJoinExpires: gog.Ptr(user.Token.Expires),
	}); err != nil {
		return errorn.ErrInsert.Wrap(err)
	}

	return c.sendEmail(user.Email, user.Token.String())
}

func (c joinByEmail) existedUser(user entity.User) error {
	if !user.Status.IsWait() {
		return errorn.ErrUnknown
	}

	if !user.Token.IsExpiredByNow() {
		return errorn.ErrUserNoConfirm
	}

	token := entity.NewTokenTTL()

	if err := c.repository.UpdateTokeJoin(user.ID, token); err != nil {
		return errorn.ErrUpdate.Wrap(err)
	}

	return c.sendEmail(user.Email, token.String())
}

func (c joinByEmail) sendEmail(emailAddr, token string) error {
	if err := c.email.Send(email.JoinData{
		Email: emailAddr,
		Token: token,
	}); err != nil {
		return errorn.ErrSendEmail.Wrap(err)
	}

	return nil
}
