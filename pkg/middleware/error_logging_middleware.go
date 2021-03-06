package middleware

import (
	"fmt"

	app_context "github.com/holocycle/holo-back/pkg/context"
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
		err := next(c)
		if err == nil {
			return nil
		}

		log := app_context.GetLog(c.Request().Context())
		log = log.With(zap.Error(err))

		detail := fmt.Sprintf("%+v", err)
		hasDetail := detail != err.Error()
		if hasDetail {
			log = log.With(zap.String("detail", detail))
		}

		if _, ok := err.(*echo.HTTPError); ok {
			log.Info("error log")
		} else {
			log.Error("error log")
			if hasDetail {
				fmt.Printf("%+v", err) // output as plain text for readablity
			}
		}
		return err
	}
}
