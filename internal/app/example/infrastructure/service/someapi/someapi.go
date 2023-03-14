package someapi

import (
	example "app-module/internal/app/example/domain"
	"app-module/pkg/errors"
	"fmt"
	"net/http"
)

type delivery struct {
	url string
}

func New(url string) example.SomeApiDelivery {
	return &delivery{url: url}
}

func (d *delivery) CheckSomeData(param string) error {
	res, err := http.DefaultClient.Get(d.url + "/" + param)
	if err != nil {
		return errors.Wrap(err, "do request")
	}

	if res.StatusCode != http.StatusOK {
		defer res.Body.Close()

		return fmt.Errorf("some data not exists")
	}

	return nil
}
