package storage

import "ca_kobercams/internal/entity"

type Repositories struct {
	UserRepository UserRepoInterface
}

type UserRepoInterface interface {
	Store(user *entity.User) error
	FindAll() (error, []entity.User)
	FindUserByUsername(username string) (error, *entity.User)
	FindUserByEmail(email string) (error, *entity.User)
}
