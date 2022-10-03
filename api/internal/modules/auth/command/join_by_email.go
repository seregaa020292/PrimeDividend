package command

import (
	"primedivident/internal/decorator"
	"primedivident/internal/modules/auth/dto"
	"primedivident/internal/modules/auth/entity"
	"primedivident/internal/modules/auth/repository"
	"primedivident/internal/modules/auth/service/email"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
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
	user, err := c.repository.FindByEmail(cmd.Email)
	if err != nil {
		return errs.NotFound.Wrap(err, errmsg.CouldNotBeFound)
	}

	if user.IsEmpty() {
		if user, err = c.newUser(cmd); err != nil {
			return err
		}
	} else {
		if user, err = c.existUser(user); err != nil {
			return err
		}
	}

	if err := c.email.Send(email.JoinData{Email: user.Email, Token: user.Token.String()}); err != nil {
		return errs.BadRequest.Wrap(err, errmsg.FailedSendMessage)
	}

	return nil
}

func (c joinByEmail) newUser(cmd Credential) (entity.User, error) {
	user, err := entity.NewUser(cmd.Email, cmd.Name, cmd.Password)
	if err != nil {
		return entity.User{}, errs.BadRequest.Wrap(err, errmsg.UnknownError)
	}

	if err := c.repository.Add(dto.ModelUserByEntity(user)); err != nil {
		return entity.User{}, errs.BadRequest.Wrap(err, errmsg.FailedAddData)
	}

	return user, nil
}

func (c joinByEmail) existUser(user entity.User) (entity.User, error) {
	if !user.Status.IsWait() {
		return entity.User{}, errs.BadRequest.New(errmsg.CheckingWhileOccurred)
	}

	if !user.Token.IsExpiredByNow() {
		return entity.User{}, errs.BadRequest.New(errmsg.CheckingWhileOccurred)
	}

	if err := c.repository.UpdateTokeJoin(user.ID, user.SetGenToken()); err != nil {
		return entity.User{}, errs.BadRequest.Wrap(err, errmsg.FailedUpdateData)
	}

	return user, nil
}
