package repository

import (
	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"primedivident/internal/models"
	"primedivident/internal/models/app/public/model"
	"primedivident/internal/models/app/public/table"
	"primedivident/internal/modules/auth/dto"
	"primedivident/internal/modules/auth/entity"
	"primedivident/pkg/db/postgres"
)

type Repository interface {
	Add(user model.Users) error
	Confirm(tokenValue uuid.UUID) error
	FindByTokenJoin(tokenValue uuid.UUID) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	HasByEmail(email string) (bool, error)
	UpdateTokeJoin(id uuid.UUID, token entity.Token) error
}

type repository struct {
	db *postgres.Postgres
}

func NewRepository(db *postgres.Postgres) Repository {
	return repository{db: db}
}

func (r repository) Add(user model.Users) error {
	stmt := table.Users.INSERT(
		table.Users.ID,
		table.Users.Email,
		table.Users.Name,
		table.Users.Password,
		table.Users.Role,
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

	return dto.EntityUserByModel(user), nil
}

func (r repository) Confirm(tokenValue uuid.UUID) error {
	stmt := table.Users.
		UPDATE(table.Users.Status, table.Users.TokenJoinValue, table.Users.TokenJoinExpires).
		SET(models.ActiveStatus, nil, nil).
		WHERE(table.Users.TokenJoinValue.EQ(jet.UUID(tokenValue)))

	_, err := stmt.Exec(r.db)

	return err
}

func (r repository) FindByEmail(email string) (entity.User, error) {
	var user model.Users

	stmt := table.Users.
		SELECT(table.Users.AllColumns).
		FROM(table.Users).
		WHERE(table.Users.Email.EQ(jet.String(email))).
		LIMIT(1)

	err := stmt.Query(r.db, &user)

	if errors.Is(err, qrm.ErrNoRows) {
		return entity.User{}, nil
	}

	if err != nil {
		return entity.User{}, err
	}

	return dto.EntityUserByModel(user), nil
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

func (r repository) UpdateTokeJoin(id uuid.UUID, token entity.Token) error {
	stmt := table.Users.
		UPDATE(table.Users.TokenJoinValue, table.Users.TokenJoinExpires).
		SET(token.Value, token.Expires).
		WHERE(table.Users.ID.EQ(jet.UUID(id)))

	_, err := stmt.Exec(r.db)

	return err
}
