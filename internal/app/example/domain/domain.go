package example

import (
	"errors"
)

// directly model
// the idea is to use universal naming to pretty imports.
// for example: no matter what entity do you use, the import of it will have appearence: user.Instance; list.Instance and so on.
type Instance struct {
	ID   int         `json:"id"`
	Test interface{} `json:"test"`
}

// Response struct for answering to clients
type Response struct {
	Data  any
	Error any
}

// Usecase behaviour
type Usecase interface {
	GetExampleData(id int) (*Instance, error)
}

// Repository behaviour
type Repository interface {
	Create(data *Instance) error
	GetOne(id int) (*Instance, error)
	GetMany(ids []int) ([]*Instance, error)
	Update(data *Instance) error
	Delete(id int) error
}

// Delivery behaviour
type Delivery interface {
	Expose()
}

// validate private funciton, contains any rules for validating/sanitazingggggg
func (ex *Instance) validate() error {
	if ex.Test == nil {
		return errors.New("test cannot be nil")
	}
	return nil
}
