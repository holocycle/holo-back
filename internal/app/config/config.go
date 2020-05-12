package config

import (
	"github.com/holocycle/holo-back/pkg/db"
	"github.com/holocycle/holo-back/pkg/logger"
	"github.com/holocycle/holo-back/pkg/middleware"
	"github.com/holocycle/holo-back/pkg/youtube"
	"github.com/jinzhu/configor"
)

const (
	envPrefix = "APP"
	fileName  = "./config/app/config.yaml"
)

type AppConfig struct {
	AppName       string `required:"true"`
	Port          string `required:"true" env:"PORT"`
	Logger        logger.Config
	DB            db.Config
	CORS          middleware.CORSConfig
	GoogleOAuth2  GoogleOAuth2Config
	YoutubeClient youtube.Config
}

func NewConfig() (*AppConfig, error) {
	c := configor.New(&configor.Config{
		ENVPrefix:            envPrefix,
		Debug:                true,
		ErrorOnUnmatchedKeys: true,
		// Verbose: true,
	})

	config := &AppConfig{}
	err := c.Load(config, fileName)
	if err != nil {
		return nil, err
	}
	return config, nil
}
