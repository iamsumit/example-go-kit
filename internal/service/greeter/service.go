package greeter

type Service interface {
	Greet(message string) (string, error)
}

type GreetService struct {
	repos Repository
}

func NewService(r Repository) Service {
	return &GreetService{
		repos: r,
	}
}

func (s *GreetService) Greet(message string) (string, error) {
	return s.repos.Greet(message)
}
