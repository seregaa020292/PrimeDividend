package model

import (
	"primedivident/internal/decorator"
	"primedivident/pkg/db/postgres"
)

type (
	instrument decorator.Model[columns]
	columns    struct {
		ID          string
		Title       string
		Description string
		CreatedAt   string
		UpdatedAt   string
	}
)

var Instrument = instrument{
	Table: "instruments",
	Columns: columns{
		ID:          "id",
		Title:       "title",
		Description: "description",
		CreatedAt:   "created_at",
		UpdatedAt:   "updated_at",
	},
}

func (m instrument) SelectAll() string {
	sql, _ := postgres.Builder.Select("*").From(m.Table).MustSql()
	return sql
}
