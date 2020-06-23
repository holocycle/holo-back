package middleware

import (
	app_context "github.com/holocycle/holo-back/pkg/context"
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
		log := app_context.GetLog(c.Request().Context())

		param, _ := c.FormParams()
		pathParam := c.ParamValues()

		req := c.Request()
		log.Info("request log",
			zap.String("method", req.Method),
			zap.String("host", req.Host),
			zap.String("path", c.Path()),
			zap.Any("param", param),
			zap.Any("pathParam", pathParam))

		return next(c)
	}
}
