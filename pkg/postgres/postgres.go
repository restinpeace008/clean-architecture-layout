package postgres

import (
	"app-module/pkg/errors"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

func New() *sql.DB {
	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
			viper.GetString("postgres.host"),
			viper.GetInt("postgres.port"),
			viper.GetString("postgres.user"),
			viper.GetString("postgres.password"),
			viper.GetString("postgres.database")))
	if err != nil {
		log.Fatal(errors.Wrap(err, "unable to connect"))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal(errors.Wrap(err, "ping db"))
	}

	return db
}
