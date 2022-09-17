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

var Assets = newAssetsTable("public", "assets", "")

type assetsTable struct {
	postgres.Table

	//Columns
	ID          postgres.ColumnString
	Amount      postgres.ColumnInteger
	Quantity    postgres.ColumnInteger
	PortfolioID postgres.ColumnString
	MarketID    postgres.ColumnString
	NotationAt  postgres.ColumnTimestampz
	CreatedAt   postgres.ColumnTimestampz
	UpdatedAt   postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type AssetsTable struct {
	assetsTable

	EXCLUDED assetsTable
}

// AS creates new AssetsTable with assigned alias
func (a AssetsTable) AS(alias string) *AssetsTable {
	return newAssetsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new AssetsTable with assigned schema name
func (a AssetsTable) FromSchema(schemaName string) *AssetsTable {
	return newAssetsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new AssetsTable with assigned table prefix
func (a AssetsTable) WithPrefix(prefix string) *AssetsTable {
	return newAssetsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new AssetsTable with assigned table suffix
func (a AssetsTable) WithSuffix(suffix string) *AssetsTable {
	return newAssetsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newAssetsTable(schemaName, tableName, alias string) *AssetsTable {
	return &AssetsTable{
		assetsTable: newAssetsTableImpl(schemaName, tableName, alias),
		EXCLUDED:    newAssetsTableImpl("", "excluded", ""),
	}
}

func newAssetsTableImpl(schemaName, tableName, alias string) assetsTable {
	var (
		IDColumn          = postgres.StringColumn("id")
		AmountColumn      = postgres.IntegerColumn("amount")
		QuantityColumn    = postgres.IntegerColumn("quantity")
		PortfolioIDColumn = postgres.StringColumn("portfolio_id")
		MarketIDColumn    = postgres.StringColumn("market_id")
		NotationAtColumn  = postgres.TimestampzColumn("notation_at")
		CreatedAtColumn   = postgres.TimestampzColumn("created_at")
		UpdatedAtColumn   = postgres.TimestampzColumn("updated_at")
		allColumns        = postgres.ColumnList{IDColumn, AmountColumn, QuantityColumn, PortfolioIDColumn, MarketIDColumn, NotationAtColumn, CreatedAtColumn, UpdatedAtColumn}
		mutableColumns    = postgres.ColumnList{AmountColumn, QuantityColumn, PortfolioIDColumn, MarketIDColumn, NotationAtColumn, CreatedAtColumn, UpdatedAtColumn}
	)

	return assetsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		Amount:      AmountColumn,
		Quantity:    QuantityColumn,
		PortfolioID: PortfolioIDColumn,
		MarketID:    MarketIDColumn,
		NotationAt:  NotationAtColumn,
		CreatedAt:   CreatedAtColumn,
		UpdatedAt:   UpdatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}