package example

import (
	example "app-module/internal/app/example/domain"
	"app-module/pkg/errors"
	"app-module/pkg/source"

	"github.com/sirupsen/logrus"
)

// Repository instance
// Define our dependencies here
type Repository struct {
	someLogger *logrus.Logger
	someSource *source.Storage
}

// New `Repository` factory
// And inject them here
func New(logger *logrus.Logger, source *source.Storage) example.Repository {
	return &Repository{
		someLogger: logger,
		someSource: source,
	}
}

// The methods below must be called only by `Usecase` of this entity.
// It's anti-pattern, when some other `Usecase` calls this.

// Create demo method. implements the `Repository` interface
func (r Repository) Create(data *example.Instance) error {
	// The way using injected dependencies
	r.someLogger.Infoln("[example] Create - not implemented!")
	// Trying to call some injected method.
	r.someLogger.Infoln("[example] Create - try to PING some source... ", r.someSource.Ping())
	return nil
}

// GetOne implements the `Repository` interface
func (r *Repository) GetOne(id int) (*example.Instance, error) {
	// Trying to call another injected method, but with error now.
	// Read about this errors here: pkg/errors.
	if _, err := r.someSource.PingWithError(); err != nil {
		return nil, errors.Wrap(err, "GetOne")
	}
	return &example.Instance{ID: 1, Test: "test"}, nil
}

// GetMany implements the `Repository` interface
func (r Repository) GetMany(ids []int) ([]*example.Instance, error) {
	r.someLogger.Infoln("[example] GetMany - not implemented!")
	r.someLogger.Infoln("[example] GetMany - try to PING some source... ", r.someSource.Ping())
	return nil, nil
}

// Update implements the `Repository` interface
func (r Repository) Update(data *example.Instance) error {
	r.someLogger.Infoln("[example] Update - not implemented!")
	r.someLogger.Infoln("[example] Update - try to PING some source... ", r.someSource.Ping())
	return nil
}

// Delete implements the `Repository` interface
func (r Repository) Delete(id int) error {
	r.someLogger.Infoln("[example] Delete - not implemented!")
	r.someLogger.Infoln("[example] Delete - try to PING some source... ", r.someSource.Ping())
	return nil
}
