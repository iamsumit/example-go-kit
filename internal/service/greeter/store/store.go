package store

type Store struct {
}

func New() *Store {
	return &Store{}
}

func (s *Store) Greet(msg string) string {
	return msg
}
