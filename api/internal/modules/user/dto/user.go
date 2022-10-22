package dto

import (
	jet "github.com/go-jet/jet/v2/postgres"

	"primedivident/internal/decorators"
	"primedivident/internal/models/app/public/table"
)

type UpdateVariadic = decorators.ColumnAssigment[any]

func NewUpdateVariadic(name *string, email *string) UpdateVariadic {
	columns := make(UpdateVariadic, 0)

	if name != nil {
		columns = append(columns, table.Users.Name.SET(jet.String(*name)))
	}
	if email != nil {
		columns = append(columns, table.Users.Email.SET(jet.String(*email)))
	}

	return columns
}
