package youtube_client

type APIURL struct {
	Video string `required:"true"`
}

type Config struct {
	APIKey string `required:"true" env:"YOUTUBE_CLIENT_API_KEY"`
	APIURL APIURL
}

type Client struct {
	APIKey string
	APIURL APIURL
}

func New(config *Config) *Client {
	return &Client{
		APIKey: config.APIKey,
		APIURL: config.APIURL,
	}
}
