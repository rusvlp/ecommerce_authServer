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

	var username string = user.Username
	var password string = user.Password
	var email string = user.Email

	print("trying to insert user with creds: ", user.Email, ", ", user.Password, ", ", user.Email)

	sqlQuery, args, err := sq.Insert("users").
		Columns("username", "password", "email").
		Values(username, password, email).
		ToSql()
	fmt.Println(sqlQuery)

	if err != nil {
		return err
	}

	_, err = repository.Database.Exec(sqlQuery, args...)

	if err != nil {
		print(err.Error())
		return err
	}
	return nil

}

func FindUserByUsername(username string) (error, *entity.User) {
	return nil, nil
}
func FindUserByEmail(email string) (error, *entity.User) {
	return nil, nil
}

func (repository *UserRepository) FindAll() (error, []entity.User) {
	return nil, nil
}

/*
Store(user *entity.User) error
FindAll() (error, []entity.User)
*/
