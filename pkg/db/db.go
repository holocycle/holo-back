package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
}

func NewDB(config *DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password))

	if err != nil {
		return nil, err
	}

	return db, nil
}
