package repository

import (
	"primedivident/internal/modules/instrument/entity"
	"primedivident/pkg/db/postgres"
)

type Repository struct {
	db *postgres.Postgres
}

func NewRepository(db *postgres.Postgres) Repository {
	return Repository{db: db}
}

func (r Repository) GetAll() (entity.Instruments, error) {
	var instruments entity.Instruments

	err := r.db.Select(&instruments, "SELECT * FROM instruments")

	return instruments, err
}
