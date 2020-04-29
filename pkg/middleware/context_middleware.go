package middleware

import (
	"github.com/holocycle/holo-back/pkg/context"
	"github.com/labstack/echo/v4"
)

type ContextMiddleware struct{}

func NewContextMiddleware() echo.MiddlewareFunc {
	m := &ContextMiddleware{}
	return m.Process
}

func (m *ContextMiddleware) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(context.NewContextFrom(c))
	}
}
