package mysql

import "database/sql"

type repositories struct{
	UserRepository *UserRepository
}

func InitRepositories(database *sql.DB) (error, *repositories){
	
	repos := &repositories{
		UserRepository:
	}
}

type Mysql struct{

}
