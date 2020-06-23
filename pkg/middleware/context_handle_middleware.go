package middleware

import (
	"context"

	"github.com/labstack/echo/v4"
)

type ContextHandler func(ctx context.Context) (context.Context, error)

type ContextHandleMiddleware struct {
	Handler ContextHandler
}

func NewContextHandleMiddleware(handler ContextHandler) echo.MiddlewareFunc {
	m := &ContextHandleMiddleware{
		Handler: handler,
	}
	return m.Process
}

func (m *ContextHandleMiddleware) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		newCtx, err := m.Handler(c.Request().Context())
		if err != nil {
			return err
		}
		c.SetRequest(c.Request().WithContext(newCtx))
		return next(c)
	}
}
