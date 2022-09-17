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

var Currencies = newCurrenciesTable("public", "currencies", "")

type currenciesTable struct {
	postgres.Table

	//Columns
	ID          postgres.ColumnString
	Title       postgres.ColumnString
	Description postgres.ColumnString
	CreatedAt   postgres.ColumnTimestampz
	UpdatedAt   postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type CurrenciesTable struct {
	currenciesTable

	EXCLUDED currenciesTable
}

// AS creates new CurrenciesTable with assigned alias
func (a CurrenciesTable) AS(alias string) *CurrenciesTable {
	return newCurrenciesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new CurrenciesTable with assigned schema name
func (a CurrenciesTable) FromSchema(schemaName string) *CurrenciesTable {
	return newCurrenciesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new CurrenciesTable with assigned table prefix
func (a CurrenciesTable) WithPrefix(prefix string) *CurrenciesTable {
	return newCurrenciesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new CurrenciesTable with assigned table suffix
func (a CurrenciesTable) WithSuffix(suffix string) *CurrenciesTable {
	return newCurrenciesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newCurrenciesTable(schemaName, tableName, alias string) *CurrenciesTable {
	return &CurrenciesTable{
		currenciesTable: newCurrenciesTableImpl(schemaName, tableName, alias),
		EXCLUDED:        newCurrenciesTableImpl("", "excluded", ""),
	}
}

func newCurrenciesTableImpl(schemaName, tableName, alias string) currenciesTable {
	var (
		IDColumn          = postgres.StringColumn("id")
		TitleColumn       = postgres.StringColumn("title")
		DescriptionColumn = postgres.StringColumn("description")
		CreatedAtColumn   = postgres.TimestampzColumn("created_at")
		UpdatedAtColumn   = postgres.TimestampzColumn("updated_at")
		allColumns        = postgres.ColumnList{IDColumn, TitleColumn, DescriptionColumn, CreatedAtColumn, UpdatedAtColumn}
		mutableColumns    = postgres.ColumnList{TitleColumn, DescriptionColumn, CreatedAtColumn, UpdatedAtColumn}
	)

	return currenciesTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		Title:       TitleColumn,
		Description: DescriptionColumn,
		CreatedAt:   CreatedAtColumn,
		UpdatedAt:   UpdatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
