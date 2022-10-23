package repository

import (
	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"

	"primedivident/internal/models/app/public/model"
	"primedivident/internal/models/app/public/table"
	"primedivident/pkg/db/postgres"
)

type Repository interface {
	GetUserAll(userID, portfolioID uuid.UUID) ([]model.Assets, error)
}

type repository struct {
	db *postgres.Postgres
}

func NewRepository(db *postgres.Postgres) Repository {
	return repository{db: db}
}

func (r repository) GetUserAll(userID, portfolioID uuid.UUID) ([]model.Assets, error) {
	var assets []model.Assets

	stmt := jet.SELECT(table.Assets.AllColumns, table.Portfolios.AllColumns).
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
