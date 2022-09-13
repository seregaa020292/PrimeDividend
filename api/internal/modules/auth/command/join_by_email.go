package command

import (
	"github.com/google/uuid"

	"primedivident/internal/decorator"
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
	user := entity.User{
		Email:     cmd.Email,
		Password:  cmd.Password,
		Confirmed: gog.Ptr(uuid.New()),
	}

	if err := c.repository.Add(user); err != nil {
		return errorn.Authorization(errorn.Message{Error: err})
	}

	if err := c.email.Send(email.ConfirmData{
		Email: user.Email,
		Token: user.Confirmed.String(),
	}); err != nil {
		return errorn.Authorization(errorn.Message{Error: err})
	}

	return nil
}
