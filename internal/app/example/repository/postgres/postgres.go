package postgres

import (
	example "app-module/internal/app/example/domain"
	"app-module/pkg/errors"
	"app-module/pkg/postgres"

	"github.com/sirupsen/logrus"
)

// Repository instance
// Define our dependencies here
type Repository struct {
	logger *logrus.Logger
	sql    *postgres.Postgres
}

// New `Repository` factory
// And inject them here
func New(logger *logrus.Logger, db *postgres.Postgres) example.Repository {
	return &Repository{
		logger: logger,
		sql:    db,
	}
}

// Create demo method. implements the `Repository` interface
func (r Repository) Create(data *example.Instance) error {
	if _, err := r.sql.DB.Exec("INSERT INTO example (name) VALUES ($1)", "tester"); err != nil {
		return errors.Wrap(err, "sql exec")
	}
	return nil
}

// GetOne implements the `Repository` interface
func (r *Repository) GetOne(id int) (*example.Instance, error) {
	if err := r.sql.DB.QueryRow("SELECT name FROM example WHERE id=$1", id).Err(); err != nil {
		return nil, errors.Wrap(err, "sql query row")
	}
	return nil, nil
}

// GetMany implements the `Repository` interface
func (r Repository) GetMany(ids []int) ([]*example.Instance, error) {
	return nil, nil
}

// Update implements the `Repository` interface
func (r Repository) Update(data *example.Instance) error {
	return nil
}

// Delete implements the `Repository` interface
func (r Repository) Delete(id int) error {
	return nil
}
