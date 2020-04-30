package middleware

import (
	"github.com/holocycle/holo-back/pkg/context"
	"github.com/labstack/echo/v4"
	echo_middleware "github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func NewResponseLoggingMiddleware() echo.MiddlewareFunc {
	handler := func(c echo.Context, reqBody, resBody []byte) {
		ctx := c.(context.Context)
		log := ctx.GetLog()

		param, _ := ctx.FormParams()
		pathParam := ctx.ParamValues()

		req := ctx.Request()
		res := ctx.Response()
		log.Info("response log",
			zap.String("method", req.Method),
			zap.String("host", req.Host),
			zap.String("path", ctx.Path()),
			zap.Any("param", param),
			zap.Any("pathParam", pathParam),
			zap.Int("status", res.Status),
			zap.String("response", string(resBody)),
		)
	}
	return echo_middleware.BodyDump(handler)
}
