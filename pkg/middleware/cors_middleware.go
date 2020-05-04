package middleware

import (
	"github.com/labstack/echo/v4"
	echo_middleware "github.com/labstack/echo/v4/middleware"
)

type CORSConfig struct {
	OriginWhitelist []string `required: "true"`
	HeaderWhitelist []string `required: "true"`
}

func NewCORSMiddleware(config *CORSConfig) echo.MiddlewareFunc {
	return echo_middleware.CORSWithConfig(echo_middleware.CORSConfig{
		AllowOrigins: config.OriginWhitelist,
		AllowHeaders: config.HeaderWhitelist,
	})
}
