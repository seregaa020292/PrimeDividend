package repository

import (
	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"

	"primedividend/api/internal/models/app/public/model"
	"primedividend/api/internal/models/app/public/table"
	"primedividend/api/pkg/db/postgres"
)

type Repository interface {
	FindById(id uuid.UUID) (model.Instruments, error)
	GetAll() ([]model.Instruments, error)
}

type repository struct {
	db *postgres.Postgres
}

func NewRepository(db *postgres.Postgres) Repository {
	return repository{db: db}
}

func (r repository) FindById(id uuid.UUID) (model.Instruments, error) {
	var instrument model.Instruments

	stmt := table.Instruments.
		SELECT(table.Instruments.AllColumns).
		FROM(table.Instruments).
		WHERE(table.Instruments.ID.EQ(jet.UUID(id))).
		LIMIT(1)

	err := stmt.Query(r.db, &instrument)

	return instrument, err
}

func (r repository) GetAll() ([]model.Instruments, error) {
	var instruments []model.Instruments

	stmt := table.Instruments.
		SELECT(table.Instruments.AllColumns).
		FROM(table.Instruments)

	err := stmt.Query(r.db, &instruments)

	return instruments, err
}
