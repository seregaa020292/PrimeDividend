package command

import (
	"primedivident/internal/decorator"
	"primedivident/internal/modules/auth/dto"
	"primedivident/internal/modules/auth/entity"
	"primedivident/internal/modules/auth/repository"
	"primedivident/internal/modules/auth/service/email"
	"primedivident/pkg/errorn"
)

type (
	Credential struct {
		Email    string
		Name     string
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

	user, err := entity.NewUser(cmd.Email, cmd.Name, cmd.Password)
	if err != nil {
		return errorn.ErrUnknown.Wrap(err)
	}

	if err := c.repository.Add(dto.ModelUserByEntity(user)); err != nil {
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
