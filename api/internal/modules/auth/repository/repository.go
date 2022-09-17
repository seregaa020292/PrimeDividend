package repository

import (
	"primedivident/internal/models/app/public/model"
	"primedivident/internal/models/app/public/table"
	"primedivident/pkg/db/postgres"
)

type Repository interface {
	Add(user model.Users) error
	Confirm(tokenValue string) error
	HasByEmail(email string) (bool, error)
}

type repository struct {
	db *postgres.Postgres
}

func NewRepository(db *postgres.Postgres) Repository {
	return repository{db: db}
}

func (r repository) Add(user model.Users) error {
	_, err := table.Users.INSERT(
		table.Users.Email,
		table.Users.Password,
		table.Users.TokenJoinValue,
		table.Users.TokenJoinExpires,
	).VALUES(
		user.Email,
		user.Password,
		user.TokenJoinValue,
		user.TokenJoinExpires,
	).Exec(r.db)

	return err
}

func (r repository) Confirm(tokenValue string) error {
	return nil
}

func (r repository) HasByEmail(email string) (bool, error) {
	return false, nil
}
