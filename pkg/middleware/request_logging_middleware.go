package middleware

import (
	"github.com/holocycle/holo-back/pkg/context"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type RequestLoggingMiddleware struct {
}

func NewRequestLoggingMiddleware() echo.MiddlewareFunc {
	m := &RequestLoggingMiddleware{}
	return m.Process
}

func (m *RequestLoggingMiddleware) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(context.Context)
		log := ctx.GetLog()

		param, _ := ctx.FormParams()
		pathParam := ctx.ParamValues()

		req := ctx.Request()
		log.Info("request log",
			zap.String("method", req.Method),
			zap.String("host", req.Host),
			zap.String("path", ctx.Path()),
			zap.Any("param", param),
			zap.Any("pathParam", pathParam))

		return next(ctx)
	}
}
