package repository

import (
	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"

	"primedivident/internal/models/app/public/model"
	"primedivident/internal/models/app/public/table"
	"primedivident/pkg/db/postgres"
)

type Repository interface {
	FindById(id uuid.UUID) (model.Currencies, error)
	GetAll() ([]model.Currencies, error)
}

type repository struct {
	db *postgres.Postgres
}

func NewRepository(db *postgres.Postgres) Repository {
	return repository{db: db}
}

func (r repository) FindById(id uuid.UUID) (model.Currencies, error) {
	var currency model.Currencies

	stmt := table.Currencies.
		SELECT(table.Currencies.AllColumns).
		FROM(table.Currencies).
		WHERE(table.Currencies.ID.EQ(jet.UUID(id))).
		LIMIT(1)

	err := stmt.Query(r.db, &currency)

	return currency, err
}

func (r repository) GetAll() ([]model.Currencies, error) {
	var currencies []model.Currencies

	stmt := table.Currencies.
		SELECT(table.Currencies.AllColumns).
		FROM(table.Currencies)

	err := stmt.Query(r.db, &currencies)

	return currencies, err
}
