package config

import (
	"database/sql"
	"fmt"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func InitDb() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=password dbname=main_db port=5432 sslmode=disable")
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected")

	return db, nil
}
