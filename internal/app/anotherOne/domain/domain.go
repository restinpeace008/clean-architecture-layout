package anotherOne

type Usecase interface {
}

type Repository interface {
	Create(data *Instance) error
	GetOne(id int) (*Instance, error)
	GetMany(ids []int) ([]*Instance, error)
	Update(data *Instance) error
	Delete(id int) error
}

type Delivery interface {
}

type Service interface {
}

type Request struct {
}

type Response struct {
	SomeData interface{} `json:"data"`
}

// Model
type Instance struct {
	ID   int    // `json:"id"`
	Type string // `json:"name"`
	Data string // `json:"data"`
}

func (r *Request) Validate() error {
	// errors.New("Some validate error", http.StatusUnprocessableEntity)
	return nil
}
