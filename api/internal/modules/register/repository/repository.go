package repository

import (
	"primedividend/api/internal/models/app/public/model"
	"primedividend/api/internal/models/app/public/table"
	"primedividend/api/pkg/db/postgres"
)

type Repository interface {
	Add(model.Registers) (model.Registers, error)
}

type repository struct {
	db *postgres.Postgres
}

func NewRepository(db *postgres.Postgres) Repository {
	return repository{db: db}
}

func (r repository) Add(newRegister model.Registers) (model.Registers, error) {
	var register model.Registers

	stmt := table.Registers.INSERT(
		table.Registers.Identify,
		table.Registers.ProviderID,
		table.Registers.MarketID,
	).
		MODEL(newRegister).
		RETURNING(table.Registers.AllColumns)

	err := stmt.Query(r.db, &register)

	return register, err
}
