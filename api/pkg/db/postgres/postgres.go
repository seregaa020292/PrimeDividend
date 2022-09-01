package postgres

import (
	"log"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"

	"primedivident/internal/config"
)

const (
	maxOpenConns    = 60
	maxIdleConns    = 30
	connMaxLifetime = 120 * time.Second
	connMaxIdleTime = 20 * time.Second
)

type Postgres struct {
	*sqlx.DB
}

func NewPostgres(config config.Postgres) *Postgres {
	log.Println("Start Postgres")

	connect, err := sqlx.Connect("pgx", config.Dsn())
	if err != nil {
		log.Fatalln(err)
	}

	connect.SetMaxOpenConns(maxOpenConns)
	connect.SetMaxIdleConns(maxIdleConns)
	connect.SetConnMaxLifetime(connMaxLifetime)
	connect.SetConnMaxIdleTime(connMaxIdleTime)

	if err := connect.Ping(); err != nil {
		log.Fatalln(err)
	}

	return &Postgres{connect}
}
