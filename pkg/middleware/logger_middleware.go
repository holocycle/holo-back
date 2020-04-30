package middleware

import (
	"github.com/holocycle/holo-back/pkg/context"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type LoggerMiddleware struct {
	Log *zap.Logger
}

func NewLoggerMiddleware(log *zap.Logger) echo.MiddlewareFunc {
	m := &LoggerMiddleware{
		Log: log,
	}
	return m.Process
}

func (m *LoggerMiddleware) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(context.Context)
		ctx.SetLog(m.Log)
		return next(ctx)
	}
}
