package example

import (
	example "app-module/internal/app/example/domain"

	"app-module/pkg/errors"
)

type Usecase struct {
	r example.Repository
}

func New(r example.Repository) example.Usecase {
	return Usecase{r: r}
}

func (uc Usecase) GetExampleData(id int) (*example.Instance, error) {
	data, err := uc.r.GetOne(0)
	if err != nil {
		return nil, errors.Wrap(err, "GetOne")
	}
	return data, nil
}
