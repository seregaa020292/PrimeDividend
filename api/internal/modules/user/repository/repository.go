package repository

import (
	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"

	"primedividend/api/internal/models"
	"primedividend/api/internal/models/app/public/model"
	"primedividend/api/internal/models/app/public/table"
	"primedividend/api/pkg/db/postgres"
)

type Repository interface {
	FindById(id uuid.UUID, status models.Status) (model.Users, error)
	Update(id uuid.UUID, update UpdatePatch) error
	Remove(id uuid.UUID) error
}

type repository struct {
	db *postgres.Postgres
}

func NewRepository(db *postgres.Postgres) Repository {
	return repository{db: db}
}

func (r repository) FindById(id uuid.UUID, status models.Status) (model.Users, error) {
	var user model.Users

	stmt := table.Users.
		SELECT(table.Users.AllColumns).
		FROM(table.Users).
		WHERE(jet.AND(
			table.Users.ID.EQ(jet.UUID(id)),
			table.Users.Status.EQ(jet.String(status.String())),
		)).
		LIMIT(1)

	err := stmt.Query(r.db, &user)

	return user, err
}

func (r repository) Remove(id uuid.UUID) error {
	stmt := table.Users.UPDATE().
		SET(table.Users.Status.SET(jet.String(models.RemoveStatus.String()))).
		WHERE(table.Users.ID.EQ(jet.UUID(id)))

	_, err := stmt.Exec(r.db)

	return err
}

func (r repository) Update(id uuid.UUID, update UpdatePatch) error {
	stmt := table.Users.UPDATE().
		SET(update.Column(), update.ColumnList()...).
		WHERE(table.Users.ID.EQ(jet.UUID(id)))

	_, err := stmt.Exec(r.db)

	return err
}
