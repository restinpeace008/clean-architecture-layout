package example

import (
	example "app-module/internal/app/example/domain"
	"app-module/pkg/source"
	"errors"

	"github.com/sirupsen/logrus"
)

type Repository struct {
	someLogger *logrus.Logger
	someSource *source.Storage
}

func New(logger *logrus.Logger, source *source.Storage) example.Repository {
	return &Repository{
		someLogger: logger,
		someSource: source,
	}
}

func (r Repository) Create(data *example.Instance) error {
	r.someLogger.Infoln("[example] Create - not implemented!")
	r.someLogger.Infoln("[example] Create - try to PING some source... ", r.someSource.Ping())
	return nil
}

func (r *Repository) GetOne(id int) (*example.Instance, error) {
	r.someLogger.Infoln("[example] GetOne - not implemented!")
	r.someLogger.Infoln("[example] GetOne - try to PING some source... ", r.someSource.Ping())
	return &example.Instance{ID: 1, Test: "test"}, errors.New("nil returned")
}

func (r Repository) GetMany(ids []int) ([]*example.Instance, error) {
	r.someLogger.Infoln("[example] GetMany - not implemented!")
	r.someLogger.Infoln("[example] GetMany - try to PING some source... ", r.someSource.Ping())
	return nil, nil
}

func (r Repository) Update(data *example.Instance) error {
	r.someLogger.Infoln("[example] Update - not implemented!")
	r.someLogger.Infoln("[example] Update - try to PING some source... ", r.someSource.Ping())
	return nil
}

func (r Repository) Delete(id int) error {
	r.someLogger.Infoln("[example] Delete - not implemented!")
	r.someLogger.Infoln("[example] Delete - try to PING some source... ", r.someSource.Ping())
	return nil
}
