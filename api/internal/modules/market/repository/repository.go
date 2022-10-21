package repository

import (
	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"

	"primedivident/internal/models/app/public/model"
	"primedivident/internal/models/app/public/table"
	"primedivident/pkg/db/postgres"
	"primedivident/pkg/paginate/cursor"
)

type Repository interface {
	GetAll(input cursor.PaginateInput) ([]model.Markets, error)
	FindById(id uuid.UUID) (model.Markets, error)
	Add(model.Markets) (model.Markets, error)
}

type repository struct {
	db *postgres.Postgres
}

func NewRepository(db *postgres.Postgres) Repository {
	return repository{db: db}
}

func (r repository) GetAll(input cursor.PaginateInput) ([]model.Markets, error) {
	var markets []model.Markets

	cursorJet := cursor.NewJet(input, table.Markets.ID, table.Markets.CreatedAt)

	stmt := cursorJet.PagingSetting(
		table.Portfolios.
			SELECT(table.Markets.AllColumns).
			FROM(table.Markets),
		nil,
	)

	err := stmt.Query(r.db, &markets)

	return markets, err
}

func (r repository) FindById(id uuid.UUID) (model.Markets, error) {
	var market model.Markets

	stmt := table.Markets.
		SELECT(table.Markets.AllColumns).
		FROM(table.Markets).
		WHERE(table.Markets.ID.EQ(jet.UUID(id))).
		LIMIT(1)

	err := stmt.Query(r.db, &market)

	return market, err
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
