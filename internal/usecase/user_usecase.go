package usecase

import "ca_kobercams/internal/storage"

type UserUseCase struct {
	repo storage.UserRepository
}

func (uuc *UserUseCase) CreateUser(username string, password string) error {
	return nil
}
