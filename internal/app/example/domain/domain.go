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

// TODO name
// Response struct for answering to clients
type Response struct {
	SomeData *Instance `json:"some-data"`
}

// TODO name
// Request struct for parsing client's data
type Request struct {
	SomeID int `json:"some-id"`
}

type DependencyMock struct {
	Args   any
	Result []any // []any{result, error} error is always at the end
}

// TestCase struct for testing functions with dependencies
type TestCase struct {
	Input  any                       // args for usecase func
	Want   map[string]DependencyMock // mocks for repository and delivery interfaces
	Result any                       // expected result (if exists) for usecase func
	Err    error                     // expected error for usecase func
}

// Usecase behaviour
type Usecase interface {
	GetExampleData(id int) (*Instance, error)
}

// TODO name
// Repository behaviour
type Repository interface {
	Create(data *Instance) error
	GetOne(id int) (*Instance, error)
	GetMany(ids []int) ([]*Instance, error)
	Update(data *Instance) error
	Delete(id int) error
}

// TODO name
// Delivery behaviour
type Delivery interface {
	Expose()
}

type SomeApiDelivery interface {
	CheckSomeData(param string) error
}

// validate private funciton, contains any rules for validating/sanitazing
func (ex *Request) Validate() error {
	if ex.SomeID == 0 {
		return errors.New("SomeID cannot be nil", http.StatusBadRequest)
	}

	return nil
}
