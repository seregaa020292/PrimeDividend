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
	hasEmail, err := c.repository.HasByEmail(cmd.Email)
	if err != nil {
		return errorn.ErrorSelect.Wrap(err)
	}
	if hasEmail {
		return errorn.ErrorExistEmail
	}

	user, err := entity.NewUser(cmd.Email, cmd.Password)
	if err != nil {
		return errorn.ErrorUnknown.Wrap(err)
	}

	if err := c.repository.Add(model.Users{
		Email:            user.Email,
		Password:         user.PassHash,
		TokenJoinValue:   gog.Ptr(user.Token.Value),
		TokenJoinExpires: gog.Ptr(user.Token.Expires),
	}); err != nil {
		return errorn.ErrorInsert.Wrap(err)
	}

	if err := c.email.Send(email.ConfirmData{
		Email: user.Email,
		Token: user.Token.String(),
	}); err != nil {
		return errorn.ErrorSendEmail.Wrap(err)
	}

	return nil
}
