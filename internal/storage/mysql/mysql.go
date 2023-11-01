package mysql

import (
	s "ca_kobercams/internal/storage"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitMysqlRepositories(db *sql.DB) (error, *s.Repositories) {

	userRepo := &UserRepository{
		Database: db,
	}

	repos := &s.Repositories{
		UserRepository: userRepo,
	}

	return nil, repos
}
