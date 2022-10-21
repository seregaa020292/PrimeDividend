package repository

import (
	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"

	"primedivident/internal/models/app/public/model"
	"primedivident/internal/models/app/public/table"
	"primedivident/pkg/db/postgres"
)

type Repository interface {
	FindById(id uuid.UUID) (model.Users, error)
}

type repository struct {
	db *postgres.Postgres
}

func NewRepository(db *postgres.Postgres) Repository {
	return repository{db: db}
}

func (r repository) FindById(id uuid.UUID) (model.Users, error) {
	var user model.Users

	stmt := table.Users.
		SELECT(table.Users.AllColumns).
		FROM(table.Users).
		WHERE(table.Users.ID.EQ(jet.UUID(id))).
		LIMIT(1)

	err := stmt.Query(r.db, &user)

	return user, err
}
