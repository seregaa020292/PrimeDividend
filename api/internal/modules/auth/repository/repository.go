package repository

import (
	"primedivident/internal/modules/auth/entity"
	"primedivident/pkg/db/postgres"
)

type Repository interface {
	Add(user entity.User) error
	Confirm(token entity.Token) error
}

type repository struct {
	db *postgres.Postgres
}

func NewRepository(db *postgres.Postgres) Repository {
	return repository{db: db}
}

func (r repository) Add(user entity.User) error {
	return nil
}

func (r repository) Confirm(token entity.Token) error {
	return nil
}
