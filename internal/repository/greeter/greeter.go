package greeter

import "errors"

type Repository struct{}

func New() *Repository {
	return &Repository{}
}

func (s *Repository) Greet(msg string) (string, error) {
	if msg == "error" {
		return "", errors.New("error while greeting")
	}

	return msg, nil
}
