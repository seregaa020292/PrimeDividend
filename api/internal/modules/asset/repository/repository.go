package repository

import (
	"context"

	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"

	"primedividend/api/internal/models/app/public/model"
	"primedividend/api/internal/models/app/public/table"
	"primedividend/api/pkg/db/postgres"
)

type Repository interface {
	GetUserAll(ctx context.Context, userID, portfolioID uuid.UUID) ([]model.Assets, error)
	HasByUser(ctx context.Context, id, userID uuid.UUID) (bool, error)
	Add(ctx context.Context, asset model.Assets) error
	Update(ctx context.Context, id uuid.UUID, patch UpdatePatch) error
	Remove(ctx context.Context, id uuid.UUID) error
}

type repository struct {
	db *postgres.Postgres
}

func NewRepository(db *postgres.Postgres) Repository {
	return repository{db: db}
}

func (r repository) GetUserAll(_ context.Context, userID, portfolioID uuid.UUID) ([]model.Assets, error) {
	var assets []model.Assets

	stmt := jet.SELECT(table.Assets.AllColumns).
		FROM(
			table.Assets.
				INNER_JOIN(table.Portfolios, table.Portfolios.ID.EQ(table.Assets.PortfolioID)),
		).
		WHERE(jet.AND(
			table.Assets.PortfolioID.EQ(jet.UUID(portfolioID)),
			table.Portfolios.UserID.EQ(jet.UUID(userID)),
		))

	err := stmt.Query(r.db, &assets)

	return assets, err
}

func (r repository) HasByUser(_ context.Context, id, userID uuid.UUID) (bool, error) {
	var dest struct {
		Exists bool
	}

	stmt := jet.SELECT(
		jet.EXISTS(
			table.Assets.SELECT(table.Assets.ID).
				FROM(
					table.Assets.
						INNER_JOIN(table.Portfolios, table.Portfolios.ID.EQ(table.Assets.PortfolioID)),
				).
				WHERE(jet.AND(
					table.Assets.ID.EQ(jet.UUID(id)),
					table.Portfolios.UserID.EQ(jet.UUID(userID)),
				)).
				LIMIT(1),
		),
	)

	err := stmt.Query(r.db, &dest)

	return dest.Exists, err
}

func (r repository) Add(ctx context.Context, asset model.Assets) error {
	stmt := table.Assets.INSERT(
		table.Assets.Amount,
		table.Assets.Quantity,
		table.Assets.PortfolioID,
		table.Assets.MarketID,
		table.Assets.NotationAt,
	).MODEL(asset)

	_, err := stmt.ExecContext(ctx, r.db)

	return err
}

func (r repository) Update(ctx context.Context, id uuid.UUID, patch UpdatePatch) error {
	stmt := table.Assets.UPDATE().
		SET(patch.Column(), patch.ColumnList()...).
		WHERE(table.Assets.ID.EQ(jet.UUID(id)))

	_, err := stmt.ExecContext(ctx, r.db)

	return err
}

func (r repository) Remove(ctx context.Context, id uuid.UUID) error {
	stmt := table.Assets.DELETE().
		WHERE(table.Assets.ID.EQ(jet.UUID(id)))

	_, err := stmt.ExecContext(ctx, r.db)

	return err
}
