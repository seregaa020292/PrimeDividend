package repository

import (
	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"

	"primedivident/internal/models/app/public/model"
	"primedivident/internal/models/app/public/table"
	"primedivident/internal/modules/portfolio/dto"
	"primedivident/pkg/db/postgres"
	"primedivident/pkg/paginate/cursor"
)

type Repository interface {
	FindById(id uuid.UUID) (model.Portfolios, error)
	FindByUserId(userID uuid.UUID) ([]model.Portfolios, error)
	Count(query model.Portfolios) (int, error)
	GetAll(input cursor.PaginateInput, query model.Portfolios) ([]model.Portfolios, error)
	Add(portfolio model.Portfolios) error
	Update(id, userID uuid.UUID, update dto.UpdateVariadic) error
	Remove(id, userID uuid.UUID) error
}

type repository struct {
	db *postgres.Postgres
}

func NewRepository(db *postgres.Postgres) Repository {
	return repository{db: db}
}

func (r repository) FindById(id uuid.UUID) (model.Portfolios, error) {
	var portfolio model.Portfolios

	stmt := table.Portfolios.
		SELECT(table.Portfolios.AllColumns).
		FROM(table.Portfolios).
		WHERE(table.Portfolios.ID.EQ(jet.UUID(id))).
		LIMIT(1)

	err := stmt.Query(r.db, &portfolio)

	return portfolio, err
}

func (r repository) FindByUserId(userID uuid.UUID) ([]model.Portfolios, error) {
	var portfolios []model.Portfolios

	stmt := table.Portfolios.
		SELECT(table.Portfolios.AllColumns).
		FROM(table.Portfolios).
		WHERE(jet.AND(
			table.Portfolios.UserID.EQ(jet.UUID(userID)),
			table.Portfolios.Active.EQ(jet.Bool(true)),
		))

	err := stmt.Query(r.db, &portfolios)

	return portfolios, err
}

func (r repository) Count(query model.Portfolios) (int, error) {
	var dest struct {
		Count int
	}

	stmt := table.Portfolios.
		SELECT(jet.COUNT(jet.STAR).AS("count")).
		FROM(table.Portfolios)

	err := stmt.Query(r.db, &dest)

	return dest.Count, err
}

func (r repository) GetAll(input cursor.PaginateInput, query model.Portfolios) ([]model.Portfolios, error) {
	var portfolios []model.Portfolios

	cursorJet := cursor.NewJet(input, table.Portfolios.ID, table.Portfolios.CreatedAt)

	stmt := cursorJet.PagingSetting(
		table.Portfolios.
			SELECT(table.Portfolios.AllColumns).
			FROM(table.Portfolios),
		table.Portfolios.Active.EQ(jet.Bool(query.Active)),
	)

	err := stmt.Query(r.db, &portfolios)

	return portfolios, err
}

func (r repository) Add(portfolio model.Portfolios) error {
	stmt := table.Portfolios.INSERT(
		table.Portfolios.Title,
		table.Portfolios.Active,
		table.Portfolios.UserID,
		table.Portfolios.CurrencyID,
	).MODEL(portfolio)

	_, err := stmt.Exec(r.db)

	return err
}

func (r repository) Update(id, userID uuid.UUID, update dto.UpdateVariadic) error {
	stmt := table.Portfolios.UPDATE().
		SET(update.Column(), update.ColumnList()...).
		WHERE(jet.AND(
			table.Portfolios.ID.EQ(jet.UUID(id)),
			table.Portfolios.UserID.EQ(jet.UUID(userID)),
		))

	_, err := stmt.Exec(r.db)

	return err
}

func (r repository) Remove(id, userID uuid.UUID) error {
	stmt := table.Portfolios.DELETE().
		WHERE(jet.AND(
			table.Portfolios.ID.EQ(jet.UUID(id)),
			table.Portfolios.UserID.EQ(jet.UUID(userID)),
		))

	_, err := stmt.Exec(r.db)

	return err
}
