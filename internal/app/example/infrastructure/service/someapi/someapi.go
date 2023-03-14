package someapi

import (
	example "app-module/internal/app/example/domain"
	"app-module/pkg/errors"
	"fmt"
	"net/http"
)

type service struct {
	url string
}

func New(url string) example.SomeService {
	return &service{url: url}
}

func (s *service) CheckSomeData(param string) error {
	res, err := http.DefaultClient.Get(s.url + "/" + param)
	if err != nil {
		return errors.Wrap(err, "do request")
	}

	if res.StatusCode != http.StatusOK {
		defer res.Body.Close()

		return fmt.Errorf("some data not exists")
	}

	return nil
}
