package repository

import (
	"primedivident/internal/models/app/public/model"
	"primedivident/internal/models/app/public/table"
	"primedivident/pkg/db/postgres"
)

type Repository interface {
	GetAll() ([]model.Currencies, error)
}

type repository struct {
	db *postgres.Postgres
}

func NewRepository(db *postgres.Postgres) Repository {
	return repository{db: db}
}

func (r repository) GetAll() ([]model.Currencies, error) {
	var currencies []model.Currencies

	stmt := table.Currencies.
		SELECT(table.Currencies.AllColumns).
		FROM(table.Currencies)

	err := stmt.Query(r.db, &currencies)

	return currencies, err
}
