package mysql

import "database/sql"

type repositories struct {
	UserRepository *UserRepository
}

func InitRepositories(database *sql.DB) (error, *repositories) {

	userRepo := &UserRepository{
		database,
	}

	repos := &repositories{
		UserRepository: userRepo,
	}

	return nil, repos
}

type Mysql struct {
}
