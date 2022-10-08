package repository

import (
	jet "github.com/go-jet/jet/v2/postgres"

	"primedivident/internal/models/app/public/model"
	"primedivident/internal/models/app/public/table"
	"primedivident/pkg/db/postgres"
)

type Repository interface {
	GetAll() ([]model.Providers, error)
	GetByTitle(title string) (model.Providers, error)
}

type repository struct {
	db *postgres.Postgres
}

func NewRepository(db *postgres.Postgres) Repository {
	return repository{db: db}
}

func (r repository) GetAll() ([]model.Providers, error) {
	var providers []model.Providers

	stmt := table.Providers.
		SELECT(table.Providers.AllColumns).
		FROM(table.Providers)

	err := stmt.Query(r.db, &providers)

	return providers, err
}

func (r repository) GetByTitle(title string) (model.Providers, error) {
	var providers model.Providers

	stmt := table.Providers.
		SELECT(table.Providers.AllColumns).
		FROM(table.Providers).
		WHERE(table.Providers.Title.EQ(jet.String(title))).
		LIMIT(1)

	err := stmt.Query(r.db, &providers)

	return providers, err
}
