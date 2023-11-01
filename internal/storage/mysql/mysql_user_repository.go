package mysql

import (
	"ca_kobercams/internal/entity"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

type UserRepository struct {
	Database *sql.DB
}

func (repository *UserRepository) Store(user *entity.User) error {

	//username := user.Username
	//password := user.Password
	//email := user.Email

	sqlQuery, _, err := sq.Insert("users").
		Columns("username", "password", "email").
		Values("aasd", "asd", "asd").
		ToSql()
	fmt.Println(sqlQuery)

	if err != nil {
		return err
	}

	_, err = repository.Database.Exec(sqlQuery)

	if err != nil {
		return err
	}
	return nil

}

func (repository *UserRepository) FindAll() (error, []entity.User) {
	return nil, nil
}

/*
Store(user *entity.User) error
FindAll() (error, []entity.User)
*/
