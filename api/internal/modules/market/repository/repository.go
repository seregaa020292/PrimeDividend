package repository

import (
	"primedivident/internal/models/app/public/model"
	"primedivident/internal/models/app/public/table"
	"primedivident/pkg/db/postgres"
)

type Repository interface {
	Add(model.Markets) (model.Markets, error)
}

type repository struct {
	db *postgres.Postgres
}

func NewRepository(db *postgres.Postgres) Repository {
	return repository{db: db}
}

func (r repository) Add(newMarket model.Markets) (model.Markets, error) {
	var market model.Markets

	stmt := table.Markets.INSERT(
		table.Markets.Title,
		table.Markets.Ticker,
		table.Markets.Content,
		table.Markets.ImageURL,
		table.Markets.CurrencyID,
		table.Markets.InstrumentID,
	).
		MODEL(newMarket).
		RETURNING(table.Markets.AllColumns)

	err := stmt.Query(r.db, &market)

	return market, err
}
