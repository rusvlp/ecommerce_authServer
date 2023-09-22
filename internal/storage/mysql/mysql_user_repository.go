package mysql

import (
	"ca_kobercams/internal/entity"
	"database/sql"
)

type UserRepository struct {
	Database *sql.DB
}

func (ur *UserRepository) Store(user *entity.User) error {
	return nil
}

func (ur *UserRepository) FindAll(user *entity.User) (error, []entity.User) {
	return nil, nil
}

/*
Store(user *entity.User) error
FindAll() (error, []entity.User)
*/
