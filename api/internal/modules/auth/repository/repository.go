package repository

import (
	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"

	"primedivident/internal/models/app/public/model"
	"primedivident/internal/models/app/public/table"
	"primedivident/internal/modules/auth/entity"
	"primedivident/pkg/db/postgres"
)

type Repository interface {
	Add(user model.Users) error
	FindByTokenJoin(tokenValue uuid.UUID) (entity.User, error)
	Confirm(tokenValue uuid.UUID) error
	HasByEmail(email string) (bool, error)
}

type repository struct {
	db *postgres.Postgres
}

func NewRepository(db *postgres.Postgres) Repository {
	return repository{db: db}
}

func (r repository) Add(user model.Users) error {
	stmt := table.Users.INSERT(
		table.Users.Email,
		table.Users.Password,
		table.Users.Status,
		table.Users.TokenJoinValue,
		table.Users.TokenJoinExpires,
	).MODEL(user)

	_, err := stmt.Exec(r.db)

	return err
}

func (r repository) FindByTokenJoin(tokenValue uuid.UUID) (entity.User, error) {
	var user model.Users

	stmt := table.Users.
		SELECT(table.Users.AllColumns).
		FROM(table.Users).
		WHERE(table.Users.TokenJoinValue.EQ(jet.UUID(tokenValue))).
		LIMIT(1)

	err := stmt.Query(r.db, &user)
	if err != nil {
		return entity.User{}, err
	}

	return entity.User{
		Email:    user.Email,
		PassHash: user.Password,
		Status:   entity.Status(user.Status),
		Token: entity.Token{
			Value:   *user.TokenJoinValue,
			Expires: *user.TokenJoinExpires,
		},
	}, nil
}

func (r repository) Confirm(tokenValue uuid.UUID) error {
	stmt := table.Users.
		UPDATE(table.Users.Status, table.Users.TokenJoinValue, table.Users.TokenJoinExpires).
		SET(entity.Active, nil, nil).
		WHERE(table.Users.TokenJoinValue.EQ(jet.UUID(tokenValue)))

	_, err := stmt.Exec(r.db)

	return err
}

func (r repository) HasByEmail(email string) (bool, error) {
	var dest struct {
		Exists bool
	}

	stmt := jet.SELECT(
		jet.EXISTS(
			table.Users.SELECT(table.Users.ID).
				FROM(table.Users).
				WHERE(table.Users.Email.EQ(jet.String(email))).
				LIMIT(1),
		),
	)

	err := stmt.Query(r.db, &dest)

	return dest.Exists, err
}
