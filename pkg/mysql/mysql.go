package mysql

import (
	"ca_kobercams/config"
	"database/sql"
	"time"
)

const (
	_defaultConnMaxLifetime = time.Minute * 3
	_defaultMaxOpenConns    = 10
	_defaultMaxIdleConns    = 10
)

type Mysql struct {
	ConnMaxLifetime time.Duration
	MaxOpenConns    int
	MaxIdleConns    int
	Database        *sql.DB
}

func InitMysql(databaseCfg config.Database, options ...Option) (error, *Mysql) {
	mysql := &Mysql{
		_defaultConnMaxLifetime,
		_defaultMaxOpenConns,
		_defaultMaxIdleConns,
		nil,
	}

	for _, o := range options {
		o(mysql)
	}

	db, err := sql.Open("mysql", databaseCfg.Username+":"+databaseCfg.Password+"@tcp("+databaseCfg.URL+")/"+databaseCfg.DBName)

	if err != nil {
		return err, nil
	}

	db.SetConnMaxLifetime(mysql.ConnMaxLifetime)
	db.SetMaxIdleConns(mysql.MaxIdleConns)
	db.SetMaxOpenConns(mysql.MaxOpenConns)
	mysql.Database = db

	return nil, mysql
}
