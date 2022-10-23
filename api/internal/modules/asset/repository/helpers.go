package repository

import (
	"time"

	jet "github.com/go-jet/jet/v2/postgres"

	"primedivident/internal/decorators"
	"primedivident/internal/models/app/public/table"
)

type UpdatePatch = decorators.ColumnAssigment[any]

func NewUpdatePatch(quantity *int32, amount *int32, notationAt *time.Time) UpdatePatch {
	columns := make(UpdatePatch, 0)

	if quantity != nil {
		columns = append(columns, table.Assets.Quantity.SET(jet.Int32(*quantity)))
	}
	if amount != nil {
		columns = append(columns, table.Assets.Amount.SET(jet.Int32(*amount)))
	}
	if notationAt != nil {
		columns = append(columns, table.Assets.NotationAt.SET(jet.TimestampzT(*notationAt)))
	}

	return columns
}
