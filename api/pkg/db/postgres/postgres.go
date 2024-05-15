package postgres

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/go-jet/jet/v2/qrm"
	_ "github.com/lib/pq"

	"primedividend/api/internal/config"
	"primedividend/api/pkg/db/transaction"
	"primedividend/api/pkg/utils/errlog"
)

const (
	maxOpenConns    = 60
	maxIdleConns    = 30
	connMaxLifetime = 120 * time.Second
	connMaxIdleTime = 20 * time.Second
)

var (
	_ qrm.Queryable  = (*Postgres)(nil)
	_ qrm.Executable = (*Postgres)(nil)
)

type Postgres struct {
	*sql.DB
}

func NewPostgres(config config.Postgres) *Postgres {
	log.Println("Start Postgres")

	connect, err := sql.Open("postgres", config.Dsn())
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

	return &Postgres{DB: connect}
}

func (p Postgres) Close() {
	log.Println("Stop Postgres")

	errlog.Println(p.DB.Close())
}

func (p Postgres) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	if tx := transaction.ExtractTx(ctx); tx != nil {
		return tx.ExecContext(ctx, query, args...)
	}
	return p.DB.ExecContext(ctx, query, args)
}

func (p Postgres) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	if tx := transaction.ExtractTx(ctx); tx != nil {
		return tx.QueryContext(ctx, query, args...)
	}
	return p.DB.QueryContext(ctx, query, args...)
}
