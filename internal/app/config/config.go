package config

import (
	"os"
	"path/filepath"

	"github.com/holocycle/holo-back/pkg/db"
	"github.com/holocycle/holo-back/pkg/logger"
	"github.com/holocycle/holo-back/pkg/middleware"
	"github.com/holocycle/holo-back/pkg/youtube"
	"github.com/jinzhu/configor"
)

const (
	envPrefix       = "APP"
	defaultFilePath = "./config/app"
	defaultFileName = "config.yaml"
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

	filePath := defaultFilePath
	if envPath := os.Getenv("CONFIGOR_PATH"); envPath != "" {
		filePath = envPath
	}

	config := &AppConfig{}
	err := c.Load(
		config,
		filepath.Join(filePath, defaultFileName),
	)
	if err != nil {
		return nil, err
	}
	return config, nil
}
