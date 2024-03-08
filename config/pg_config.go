package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PgConfig struct {
	UserName string
	Password string
	DBName   string
}

func NewPgConfig(userName, password, dbName string) *PgConfig {
	return &PgConfig{
		UserName: userName,
		Password: password,
		DBName:   dbName,
	}
}

func (pg *PgConfig) GetConnStr() string {
	return "user=" + pg.UserName + " password=" + pg.Password + " dbname=" + pg.DBName + " sslmode=disable"
}

func (pg *PgConfig) Init() (*sql.DB, error) {
	db, err := sql.Open("postgres", pg.GetConnStr())
	if err != nil {
		return nil, err
	}

	return db, nil
}
