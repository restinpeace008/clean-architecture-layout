package example

import (
	"net/http"

	example "app-module/internal/app/example/domain"
	"app-module/pkg/errors"
)

// `usecase` instance
type usecase struct {
	// Define dependecies
	r example.Repository
}

// New `usecase` factory
func New(r example.Repository) example.Usecase {
	// Inject dependencies
	return &usecase{r: r}
}

// GetExampleData demo method. Implements `Usecase` interface
func (uc *usecase) GetExampleData(id int) (*example.Instance, error) {
	// Go to `Repository` layer
	data, err := uc.r.GetOne(id)
	if err != nil {
		return nil, errors.Wrap(err, "GetExampleData", http.StatusInternalServerError)
	}

	return data, nil
}
