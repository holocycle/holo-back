package youtube

import "github.com/holocycle/holo-back/pkg/model"

type Config struct {
	APIKey string `required:"true" env:"YOUTUBE_CLIENT_API_KEY"`
	APIURL struct {
		Video string `required:"true"`
	}
}

type Client interface {
	GetVideo(videoID string) (*model.Video, error)
}

type ClientImpl struct {
	Config
}

func New(config *Config) Client {
	return &ClientImpl{*config}
}
