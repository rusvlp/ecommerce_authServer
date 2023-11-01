package mysql

import "time"

type Option func(mysql *Mysql)

func SetConnMaxLifeTime(d time.Duration) Option {
	return func(mysql *Mysql) {
		mysql.ConnMaxLifetime = d
	}
}

func SetMaxOpenConns(c int) Option {
	return func(mysql *Mysql) {
		mysql.MaxOpenConns = c
	}
}

func SetMaxIdleConns(i int) Option {
	return func(mysql *Mysql) {
		mysql.MaxIdleConns = i
	}
}
