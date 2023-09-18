package usecase

type User interface {
	CreateUser(username string, password string) error
}
