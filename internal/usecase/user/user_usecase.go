package user

import (
	"ca_kobercams/internal/entity"
	"ca_kobercams/internal/storage"
	"ca_kobercams/internal/usecase/dto"
	"errors"
)

type User struct {
	Repo storage.UserRepoInterface
}

func (uuc *User) CreateUser(user *entity.User, passwordConf string) error {

	if user.Username == "" {
		return errors.New("Username must not be empty")
	}

	if user.Password == "" {
		return errors.New("Password must not be empty")
	}

	if user.Email == "" {
		return errors.New("Email must not be empty")
	}

	if user.Password != passwordConf {
		return errors.New("Password and password confirmation does not match")
	}

	userEnt := &entity.User{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	}

	err := uuc.Repo.Store(userEnt)
	return err
}

func (uuc *User) FindAll() (error, []dto.GetUserDTO) {
	err, users := uuc.Repo.FindAll()

	if err != nil {
		return err, nil
	}

	dtoUsers := make([]dto.GetUserDTO, 1)

	for _, u := range users {
		dtoUsers = append(dtoUsers, dto.GetUserDTO{
			Id:       u.Id,
			Username: u.Username,
			Email:    u.Email,
		})
	}

	return nil, dtoUsers
}
