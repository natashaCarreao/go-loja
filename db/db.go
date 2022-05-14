package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaBanco() *sql.DB {
	conexao := "user=postgres dbname=postgres password=postgres host=127.0.0.1 port=5432 sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}
