package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DBConfig struct {
	URL string `required:"true" env:"DATABASE_URL"`
}

func NewDB(dbConfig *DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", dbConfig.URL)

	if err != nil {
		return nil, err
	}

	return db, nil
}
