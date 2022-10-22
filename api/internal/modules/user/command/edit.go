package command

import (
	"github.com/google/uuid"

	"primedivident/internal/decorators"
	"primedivident/internal/modules/user/dto"
	"primedivident/internal/modules/user/repository"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
)

type (
	PayloadEdit struct {
		UserID uuid.UUID

		Name  *string
		Email *string
	}
	Edit decorators.CommandHandler[PayloadEdit]
)

type edit struct {
	repository repository.Repository
}

func NewEdit(
	repository repository.Repository,
) Edit {
	return edit{
		repository: repository,
	}
}

func (c edit) Exec(cmd PayloadEdit) error {
	if err := c.repository.Update(cmd.UserID, dto.NewUpdateVariadic(
		cmd.Name,
		cmd.Email,
	)); err != nil {
		return errs.BadRequest.Wrap(err, errmsg.FailedUpdateData)
	}

	return nil
}
