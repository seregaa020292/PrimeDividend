package repository

import (
	"primedivident/internal/models/app/public/model"
	"primedivident/internal/models/app/public/table"
	"primedivident/pkg/db/postgres"
)

type Repository interface {
	GetAll() ([]model.Instruments, error)
}

type repository struct {
	db *postgres.Postgres
}

func NewRepository(db *postgres.Postgres) Repository {
	return repository{db: db}
}

func (r repository) GetAll() ([]model.Instruments, error) {
	var instruments []model.Instruments

	stmt := table.Instruments.
		SELECT(table.Instruments.AllColumns).
		FROM(table.Instruments)

	err := stmt.Query(r.db, &instruments)

	return instruments, err
}
