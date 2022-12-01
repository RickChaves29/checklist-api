package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://postgres:12345@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		println("error when try open database")
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		println("error when try connection on database")
	}

	println("connection on database is good")

	return db, err
}
