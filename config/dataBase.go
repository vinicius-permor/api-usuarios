package config

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// func Conn() sera exportada para a o pacote main para iniciar a conexao com o banco de dados
func Conn() (*sql.DB, error) {
	urlConn := "marcus:marcus2025@/users?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", urlConn)
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
