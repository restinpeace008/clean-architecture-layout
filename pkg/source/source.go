package source

type Storage struct {
	data string
}

func New() *Storage {
	return &Storage{data: "turned on"}
}

func (s *Storage) Ping() string {
	return "Pong!"
}
