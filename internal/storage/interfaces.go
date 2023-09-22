package storage

import "ca_kobercams/internal/entity"

type UserRepoInterface interface {
	Store(user *entity.User) error
	FindAll() (error, []entity.User)
}
