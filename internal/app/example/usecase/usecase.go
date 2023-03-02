package example

import (
	example "app-module/internal/app/example/domain"

	"app-module/pkg/errors"
)

// `Usecase` instance
type Usecase struct {
	// Define dependecies
	r example.Repository
}

// New `Usecase` factory
func New(r example.Repository) example.Usecase {
	// Inject dependencies
	return Usecase{r: r}
}

// GetExampleData demo method. Implements `Usecase` interface
func (uc Usecase) GetExampleData(id int) (*example.Instance, error) {
	// Go to `Repository` layer
	data, err := uc.r.GetOne(0)
	if err != nil {
		return nil, errors.Wrap(err, "GetExampleData")
	}

	return data, nil
}
