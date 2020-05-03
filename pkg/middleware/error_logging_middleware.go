package middleware

import (
	"github.com/holocycle/holo-back/pkg/context"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type ErrorLoggingMiddleware struct {
}

func NewErrorLoggingMiddleware() echo.MiddlewareFunc {
	m := &ErrorLoggingMiddleware{}
	return m.Process
}

func (m *ErrorLoggingMiddleware) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(context.Context)
		log := ctx.GetLog()

		err := next(ctx)
		if err != nil {
			log.Error("error log", zap.Error(err))
		}
		return err
	}
}
