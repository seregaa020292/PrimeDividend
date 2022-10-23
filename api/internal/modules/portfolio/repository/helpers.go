package repository

import (
	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"

	"primedivident/internal/decorators"
	"primedivident/internal/models/app/public/table"
)

type FilterGetAll struct {
	Active *bool
}

func (f FilterGetAll) Condition() jet.BoolExpression {
	condition := jet.Bool(true)

	if f.Active != nil {
		condition = condition.AND(table.Portfolios.Active.EQ(jet.Bool(*f.Active)))
	}

	return condition
}

type UpdatePatch = decorators.ColumnAssigment[any]

func NewUpdatePatch(
	title *string,
	currencyID *uuid.UUID,
	active *bool,
) UpdatePatch {
	columns := make(UpdatePatch, 0)

	if title != nil {
		columns = append(columns, table.Portfolios.Title.SET(jet.String(*title)))
	}
	if currencyID != nil {
		columns = append(columns, table.Portfolios.CurrencyID.SET(jet.UUID(*currencyID)))
	}
	if active != nil {
		columns = append(columns, table.Portfolios.Active.SET(jet.Bool(*active)))
	}

	return columns
}
