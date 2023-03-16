package postgres

import (
	example "app-module/internal/app/example/domain"
	"app-module/pkg/errors"
	"database/sql"

	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// Repository instance
// Define our dependencies here
type repository struct {
	logger *logrus.Logger
	db     *sql.DB
}

// New `repository` factory
// And inject them here
func New(logger *logrus.Logger, db *sql.DB) example.DBRepository {
	return &repository{
		logger: logger,
		db:     db,
	}
}

// Create demo method. implements the `Repository` interface
func (r *repository) Create(data *example.Instance) error {
	if _, err := r.db.Exec("INSERT INTO example (name) VALUES ($1)", data.Test); err != nil {
		return errors.Wrap(err, "sql exec")
	}
	return nil
}

// GetOne implements the `Repository` interface
func (r *repository) GetOne(id int) (*example.Instance, error) {
	var name string

	if err := r.db.QueryRow("SELECT name FROM example WHERE id = $1", id).Scan(&name); err != nil {
		return nil, errors.Wrap(err, "sql query row")
	}

	return &example.Instance{
		ID:   id,
		Test: name,
	}, nil
}

// GetMany implements the `Repository` interface
func (r *repository) GetMany(ids []int) ([]*example.Instance, error) {
	rows, err := r.db.Query("SELECT id, name FROM example WHERE id = ANY($1)", pq.Array(ids))
	if err != nil {
		return nil, errors.Wrap(err, "sql query")
	}

	defer rows.Close()

	result := make([]*example.Instance, 0, len(ids))

	for rows.Next() {
		var (
			id   int
			name string
		)

		if err := rows.Scan(&id, &name); err != nil {
			return nil, errors.Wrap(err, "rows scan")
		}

		result = append(result, &example.Instance{
			ID:   id,
			Test: name,
		})
	}

	return result, nil
}

// Update implements the `Repository` interface
func (r *repository) Update(data *example.Instance) error {
	if _, err := r.db.Exec("UPDATE example SET name = $2 WHERE id = $1", data.ID, data.Test); err != nil {
		return errors.Wrap(err, "sql exec")
	}
	return nil
}

// Delete implements the `Repository` interface
func (r *repository) Delete(id int) error {
	if _, err := r.db.Exec("DELETE FROM example WHERE id = $1", id); err != nil {
		return errors.Wrap(err, "sql exec")
	}
	return nil
}
