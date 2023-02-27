package example

import (
	"errors"
)

type Instance struct {
	ID   int         `json:"id"`
	Test interface{} `json:"test"`
}

type Response struct {
	Data  any
	Error any
}

type Usecase interface {
	GetExampleData(id int) (*Instance, error)
}

type Repository interface {
	Create(data *Instance) error
	GetOne(id int) (*Instance, error)
	GetMany(ids []int) ([]*Instance, error)
	Update(data *Instance) error
	Delete(id int) error
}

type Delivery interface {
	Expose()
}

func (ex *Instance) validate() error {
	if ex.Test == nil {
		return errors.New("test cannot be nil")
	}
	return nil
}
