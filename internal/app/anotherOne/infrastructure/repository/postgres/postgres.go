package postgres

import (
	anotherOne "app-module/internal/app/anotherOne/domain"
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
func New(logger *logrus.Logger, db *sql.DB) anotherOne.Repository {
	return &repository{
		logger: logger,
		db:     db,
	}
}

// Create demo method. implements the `Repository` interface
func (r *repository) Create(model *anotherOne.Instance) error {
	if _, err := r.db.Exec("INSERT INTO another_one (type, data) VALUES ($1, $2)", model.Type, model.Data); err != nil {
		return errors.Wrap(err, "anotherOne.Create")
	}
	return nil
}

// GetOne implements the `Repository` interface
func (r *repository) GetOne(id int) (*anotherOne.Instance, error) {
	var result *anotherOne.Instance

	if err := r.db.QueryRow("SELECT type, data FROM another_one WHERE id = $1", id).Scan(&result.Type, &result.Data); err != nil {
		return nil, errors.Wrap(err, "anotherOne.GetOne")
	}

	return result, nil
}

// GetMany implements the `Repository` interface
func (r *repository) GetMany(ids []int) ([]*anotherOne.Instance, error) {
	rows, err := r.db.Query("SELECT id, type, data FROM another_one WHERE id = ANY($1)", pq.Array(ids))
	if err != nil {
		return nil, errors.Wrap(err, "anotherOne.GetMany: do query")
	}

	defer rows.Close()

	result := make([]*anotherOne.Instance, 0, len(ids))

	for rows.Next() {
		var (
			id    int
			_type string
			data  string
		)

		if err := rows.Scan(&id, &_type, &data); err != nil {
			return nil, errors.Wrap(err, "rows scan")
		}

		result = append(result, &anotherOne.Instance{
			ID:   id,
			Type: _type,
			Data: data,
		})
	}

	return result, nil
}

// Update implements the `Repository` interface
func (r *repository) Update(model *anotherOne.Instance) error {
	if _, err := r.db.Exec("UPDATE another_one SET data = $2 WHERE id = $1", model.ID, model.Data); err != nil {
		return errors.Wrap(err, "anotherOne.Update")
	}
	return nil
}

// Delete implements the `Repository` interface
func (r *repository) Delete(id int) error {
	if _, err := r.db.Exec("DELETE FROM another_one WHERE id = $1", id); err != nil {
		return errors.Wrap(err, "anotherOne.Delete")
	}
	return nil
}
