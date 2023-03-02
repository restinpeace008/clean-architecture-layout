package source

import "app-module/pkg/errors"

type Storage struct {
	data string
}

func New() *Storage {
	return &Storage{data: "turned on"}
}

func (s *Storage) Ping() string {
	return "Pong!"
}

func (s *Storage) PingWithError() (string, error) {
	return "", errors.New("sql: no rows")
}
