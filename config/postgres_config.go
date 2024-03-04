package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb(dns string) (*gorm.DB, error) {
	// gorm connection to postgres
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db, nil
}
