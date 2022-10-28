package repository

import (
	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"

	"primedividend/api/internal/models/app/public/model"
	"primedividend/api/internal/models/app/public/table"
	"primedividend/api/pkg/db/postgres"
)

type Repository interface {
	FindById(id uuid.UUID) (model.Providers, error)
	GetAll() ([]model.Providers, error)
	GetByTitle(title string) (model.Providers, error)
}

type repository struct {
	db *postgres.Postgres
}

func NewRepository(db *postgres.Postgres) Repository {
	return repository{db: db}
}

func (r repository) FindById(id uuid.UUID) (model.Providers, error) {
	var provider model.Providers

	stmt := table.Providers.
		SELECT(table.Providers.AllColumns).
		FROM(table.Providers).
		WHERE(table.Providers.ID.EQ(jet.UUID(id))).
		LIMIT(1)

	err := stmt.Query(r.db, &provider)

	return provider, err
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
	var provider model.Providers

	stmt := table.Providers.
		SELECT(table.Providers.AllColumns).
		FROM(table.Providers).
		WHERE(table.Providers.Title.EQ(jet.String(title))).
		LIMIT(1)

	err := stmt.Query(r.db, &provider)

	return provider, err
}
