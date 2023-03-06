package example

import (
	"app-module/pkg/errors"
	"net/http"
)

// directly model
// the idea is to use universal naming to pretty imports.
// for example: no matter what entity do you use, the import of it will have appearence: user.Instance; list.Instance and so on.
type Instance struct {
	ID   int `json:"id"`
	Test any `json:"test"`
}

// Response struct for answering to clients
type Response struct {
	SomeData *Instance `json:"some-data"`
}

// Request struct for parsing client's data
type Request struct {
	SomeID int `json:"some-id"`
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

// validate private funciton, contains any rules for validating/sanitazing
func (ex *Request) Validate() error {
	if ex.SomeID == 0 {
		return errors.New("SomeID cannot be nil", http.StatusBadRequest)
	}

	return nil
}
