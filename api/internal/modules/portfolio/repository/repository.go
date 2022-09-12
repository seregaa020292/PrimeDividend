package repository

import (
	"time"

	"github.com/google/uuid"

	"primedivident/internal/modules/portfolio/entity"
	"primedivident/pkg/db/postgres"
)

type Repository interface {
	FindById(id uuid.UUID) (entity.Portfolio, error)
}

type repository struct {
	db *postgres.Postgres
}

func NewRepository(db *postgres.Postgres) Repository {
	return repository{db: db}
}

func (r repository) FindById(id uuid.UUID) (entity.Portfolio, error) {
	return entity.Portfolio{
		ID:        id,
		CreatedAt: time.Now(),
	}, nil
}
