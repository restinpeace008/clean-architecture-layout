package example

import (
	"net/http"

	example "app-module/internal/app/example/domain"
	"app-module/pkg/errors"
)

// `usecase` instance
type usecase struct {
	// Define dependecies
	r   example.Repository
	sad example.SomeApiDelivery
}

// New `usecase` factory
func New(r example.Repository, sad example.SomeApiDelivery) example.Usecase {
	// Inject dependencies
	return &usecase{
		r:   r,
		sad: sad,
	}
}

// GetExampleData demo method. Implements `Usecase` interface
func (uc *usecase) GetExampleData(id int) (*example.Instance, error) {
	// Go to `Repository` layer
	data, err := uc.r.GetOne(id)
	if err != nil {
		return nil, errors.Wrap(err, "GetExampleData", http.StatusInternalServerError)
	}

	if param, ok := data.Test.(string); ok {
		if err = uc.sad.CheckSomeData(param); err != nil {
			return nil, errors.Wrap(err, "GetExampleData", http.StatusBadRequest)
		}
	}

	return data, nil
}
