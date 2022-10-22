package command

import (
	"github.com/google/uuid"

	"primedivident/internal/decorators"
	"primedivident/internal/modules/user/repository"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
)

type (
	PayloadRemove = uuid.UUID
	Remove        decorators.CommandHandler[PayloadRemove]
)

type remove struct {
	repository repository.Repository
}

func NewRemove(
	repository repository.Repository,
) Remove {
	return remove{
		repository: repository,
	}
}

func (c remove) Exec(cmd PayloadRemove) error {
	if err := c.repository.Remove(cmd); err != nil {
		return errs.BadRequest.Wrap(err, errmsg.FailedDeleteData)
	}

	return nil
}
