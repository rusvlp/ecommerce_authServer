package mysql

import (
	"ca_kobercams/internal/entity"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

type userRepository struct {
	database    *sql.DB
	stmtCache   *sq.StmtCache
	stmtBuilder *sq.StatementBuilderType
}

func NewUserRepository(db *sql.DB) (error, *userRepository) {

	stmtCache := sq.NewStmtCache(db)
	stmtBuilder := sq.StatementBuilder.RunWith(stmtCache)

	userRepo := &userRepository{
		database:    db,
		stmtCache:   stmtCache,
		stmtBuilder: &stmtBuilder,
	}

	return nil, userRepo
}

func (repository *userRepository) Store(user *entity.User) error {

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

	_, err = repository.database.Exec(sqlQuery, args...)

	if err != nil {
		print(err.Error())
		return err
	}
	return nil

}

func (repository *userRepository) FindUserByUsername(username string) (error, *entity.User) {
	//dbCache := sq.NewStmtCache(repository.database)

	user := repository.stmtBuilder.Select("*").From("users").Where("username")

	return nil, nil
}
func (repository *userRepository) FindUserByEmail(email string) (error, *entity.User) {
	return nil, nil
}

func (repository *userRepository) FindAll() (error, []entity.User) {
	return nil, nil
}

/*
Store(user *entity.User) error
FindAll() (error, []entity.User)
*/
