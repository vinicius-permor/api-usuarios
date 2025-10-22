package config

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Conn() (*sqlx.DB, error) {
	urlConn := "marcus:marcus2025@/users?charset=utf8&parseTime=True&loc=Local"
	db, err := sqlx.Connect("mysql", urlConn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	return db, nil
}
