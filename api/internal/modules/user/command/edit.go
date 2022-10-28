package command

import (
	"github.com/google/uuid"

	"primedividend/api/internal/decorators"
	"primedividend/api/internal/modules/user/repository"
	"primedividend/api/pkg/errs"
	"primedividend/api/pkg/errs/errmsg"
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
	if err := c.repository.Update(cmd.UserID, repository.NewUpdatePatch(
		cmd.Name,
		cmd.Email,
	)); err != nil {
		return errs.BadRequest.Wrap(err, errmsg.FailedUpdateData)
	}

	return nil
}
