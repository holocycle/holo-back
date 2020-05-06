package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // disable-lint
)

type Config struct {
	URL string `required:"true" env:"DATABASE_URL"`
}

func NewDB(dbConfig *Config) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", dbConfig.URL)
	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	return db, nil
}
