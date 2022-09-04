package repository

import (
	"primedivident/internal/modules/instrument/entity"
	"primedivident/pkg/db/postgres"
)

type Repository interface {
	GetAll() (entity.Instruments, error)
}

type repository struct {
	db *postgres.Postgres
}

func NewRepository(db *postgres.Postgres) Repository {
	return repository{db: db}
}

func (r repository) GetAll() (entity.Instruments, error) {
	var instruments entity.Instruments

	err := r.db.Select(&instruments, "SELECT * FROM instruments")

	return instruments, err
}
