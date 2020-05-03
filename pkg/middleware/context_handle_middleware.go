package middleware

import (
	"github.com/holocycle/holo-back/pkg/context"
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
		ctx := c.(context.Context)

		newCtx, err := m.Handler(ctx)
		if err != nil {
			return err
		}
		return next(newCtx)
	}
}
