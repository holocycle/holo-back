package youtube

import "github.com/holocycle/holo-back/pkg/model"

type Config struct {
	APIKey string `required:"true" env:"YOUTUBE_CLIENT_API_KEY"`
	APIURL struct {
		Video   string `required:"true"`
		Channel string `required:"true"`
	}
}

type Client interface {
	GetVideo(videoID string) (*model.Video, error)
	GetChannel(channelID string) (*model.Channel, error)
	ListChannels(channelIDs []string) ([]*model.Channel, error)
}

type ClientImpl struct {
	Config
}

func New(config *Config) Client {
	return &ClientImpl{*config}
}
