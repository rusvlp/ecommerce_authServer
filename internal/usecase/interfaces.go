package usecase

import (
	"ca_kobercams/internal/entity"
	"ca_kobercams/internal/usecase/dto"
)

type UserInterface interface {
	CreateUser(request *entity.User, passwordConf string) error
	FindAll() (error, []dto.GetUserDTO)
}
