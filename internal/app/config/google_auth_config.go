package config

type GoogleOAuth2Config struct {
	GoogleAuthURL      string `required:"true"`
	GoogleTokenURL     string `required:"true"`
	GoogleTokenInfoURL string `required:"true"`
	ClientID           string `required:"true"`
	ClientSecret       string `required:"true" env:"GOOGLE_OAUTH2_CLIENT_SECRET"`
	ClientRedirectURL  string `required:"true"`
	Scope              string `required:"true"`
}
