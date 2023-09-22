package mock

import "ca_kobercams/internal/entity"

type MockUserRepository struct {
	Users []entity.User
}

func (repository *MockUserRepository) Store(user *entity.User) error {
	repository.Users = append(repository.Users, *user)
	user.Id = int64(len(repository.Users) - 1)
	return nil
}

func (repository *MockUserRepository) FindAll() (error, []entity.User) {
	return nil, repository.Users
}
