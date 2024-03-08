package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormConfig struct {
	host     string
	user     string
	password string
	dbname   string
}

func NewGormConfig(host, user, password, dbname string) *GormConfig {
	return &GormConfig{
		host:     host,
		user:     user,
		password: password,
		dbname:   dbname,
	}
}

func (g *GormConfig) GetDns() string {
	return "host=" + g.host + " user=" + g.user + " password=" + g.password + " dbname=" + g.dbname + " port=5432 sslmode=disable TimeZone=Asia/Bangkok"
}

func (g *GormConfig) Init() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(g.GetDns()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	return db, nil
}
