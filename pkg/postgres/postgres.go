package postgres

import (
	"app-module/pkg/errors"
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/lib/pq"
	"github.com/spf13/viper"
)

type Postgres struct {
	DB *sql.DB
}

func New() *Postgres {
	connString, err := pq.ParseURL(viper.GetString("postgres.dsn"))
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

	if err := initTables(db); err != nil {
		log.Fatal(errors.Wrap(err, "init tables"))
	}

	return &Postgres{DB: db}
}

func initTables(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return errors.Wrap(err, "begin transaction")
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	createExampleTable := `CREATE TABLE IF NOT EXISTS example
	(
		id                 serial       NOT NULL,
    	name               varchar(255)  NOT NULL,
    	CONSTRAINT PK_example_id PRIMARY KEY (id)
	);
	`

	_, err = tx.Exec(createExampleTable)
	if err != nil {
		return errors.Wrap(err, "exec")
	}

	return nil
}
