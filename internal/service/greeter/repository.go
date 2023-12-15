package greeter

type Repository interface {
	Greet(string) (string, error)
}
