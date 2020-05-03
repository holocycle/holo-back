package config

import (
	"github.com/holocycle/holo-back/pkg/db"
	"github.com/holocycle/holo-back/pkg/logger"
	"github.com/jinzhu/configor"
)

const (
	envPrefix = "APP"
	fileName  = "./config/app/config.yaml"
)

type AppConfig struct {
	AppName      string `required:"true"`
	Port         string `required:"true" env:"PORT"`
	Logger       logger.LoggerConfig
	DB           db.DBConfig
	GoogleOAuth2 GoogleOAuth2Config
}

func NewConfig() (*AppConfig, error) {
	c := configor.New(&configor.Config{
		ENVPrefix:            envPrefix,
		Verbose:              true,
		ErrorOnUnmatchedKeys: true,
	})

	config := &AppConfig{}
	err := c.Load(config, fileName)
	if err != nil {
		return nil, err
	}
	return config, nil
}
