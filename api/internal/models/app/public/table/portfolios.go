//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Portfolios = newPortfoliosTable("public", "portfolios", "")

type portfoliosTable struct {
	postgres.Table

	//Columns
	ID         postgres.ColumnString
	Title      postgres.ColumnString
	Active     postgres.ColumnBool
	UserID     postgres.ColumnString
	CurrencyID postgres.ColumnString
	CreatedAt  postgres.ColumnTimestampz
	UpdatedAt  postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PortfoliosTable struct {
	portfoliosTable

	EXCLUDED portfoliosTable
}

// AS creates new PortfoliosTable with assigned alias
func (a PortfoliosTable) AS(alias string) *PortfoliosTable {
	return newPortfoliosTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new PortfoliosTable with assigned schema name
func (a PortfoliosTable) FromSchema(schemaName string) *PortfoliosTable {
	return newPortfoliosTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new PortfoliosTable with assigned table prefix
func (a PortfoliosTable) WithPrefix(prefix string) *PortfoliosTable {
	return newPortfoliosTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new PortfoliosTable with assigned table suffix
func (a PortfoliosTable) WithSuffix(suffix string) *PortfoliosTable {
	return newPortfoliosTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newPortfoliosTable(schemaName, tableName, alias string) *PortfoliosTable {
	return &PortfoliosTable{
		portfoliosTable: newPortfoliosTableImpl(schemaName, tableName, alias),
		EXCLUDED:        newPortfoliosTableImpl("", "excluded", ""),
	}
}

func newPortfoliosTableImpl(schemaName, tableName, alias string) portfoliosTable {
	var (
		IDColumn         = postgres.StringColumn("id")
		TitleColumn      = postgres.StringColumn("title")
		ActiveColumn     = postgres.BoolColumn("active")
		UserIDColumn     = postgres.StringColumn("user_id")
		CurrencyIDColumn = postgres.StringColumn("currency_id")
		CreatedAtColumn  = postgres.TimestampzColumn("created_at")
		UpdatedAtColumn  = postgres.TimestampzColumn("updated_at")
		allColumns       = postgres.ColumnList{IDColumn, TitleColumn, ActiveColumn, UserIDColumn, CurrencyIDColumn, CreatedAtColumn, UpdatedAtColumn}
		mutableColumns   = postgres.ColumnList{TitleColumn, ActiveColumn, UserIDColumn, CurrencyIDColumn, CreatedAtColumn, UpdatedAtColumn}
	)

	return portfoliosTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:         IDColumn,
		Title:      TitleColumn,
		Active:     ActiveColumn,
		UserID:     UserIDColumn,
		CurrencyID: CurrencyIDColumn,
		CreatedAt:  CreatedAtColumn,
		UpdatedAt:  UpdatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
