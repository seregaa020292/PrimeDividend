package dto

import (
	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"

	"primedivident/internal/decorators"
	"primedivident/internal/models/app/public/table"
)

type PortfolioVariadic = decorators.ColumnAssigment[any]

func NewPortfolioVariadic(
	title *string,
	currencyID *uuid.UUID,
	active *bool,
) PortfolioVariadic {
	columns := make(PortfolioVariadic, 0)

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
