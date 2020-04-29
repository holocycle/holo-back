package config

import (
	"github.com/jinzhu/configor"
)

const (
	envPrefix = "APP"
	fileName  = "./config/app/config.yaml"
)

type AppConfig struct {
	AppName     string `required:"true"`
	Port        string `required:"true" env:"PORT"`
	DatabaseURL string `required:"true" env:"DATABASE_URL"`
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
