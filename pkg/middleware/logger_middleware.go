package middleware

import (
	"context"

	app_context "github.com/holocycle/holo-back/pkg/context"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type LoggerMiddleware struct {
	Log *zap.Logger
}

func NewLoggerMiddleware(log *zap.Logger) echo.MiddlewareFunc {
	return NewContextHandleMiddleware(func(ctx context.Context) (context.Context, error) {
		return app_context.SetLog(ctx, log), nil
	})
}
