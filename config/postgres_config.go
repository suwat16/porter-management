package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func InitDb() (*gorm.DB, error) {
	dns := "host=localhost user=postgres password=password dbname=main_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected")

	return db, nil
}
