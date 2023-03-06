package postgres

import (
	"app-module/pkg/errors"
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/lib/pq"
)

type Postgres struct {
	DB *sql.DB
}

func New() *Postgres {
	return new(Postgres) // FIXME

	connString, err := pq.ParseURL("POSTGRES_DSN_FROM_CONFIG")
	if err != nil {
		log.Fatal(errors.Wrap(err, "invalid dsn"))
	}

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(errors.Wrap(err, "unable to connect"))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal(errors.Wrap(err, "ping db"))
	}

	return &Postgres{DB: db}
}

// TODO migrations and other
