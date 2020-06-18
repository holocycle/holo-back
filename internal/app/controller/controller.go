package controller

import (
	app_middleware "github.com/holocycle/holo-back/internal/app/middleware"
	"github.com/holocycle/holo-back/pkg/context"
	"github.com/labstack/echo/v4"
)

type controller func(c context.Context) error

func get(e *echo.Echo, path string, controller controller) {
	e.GET(path, wrapController(controller),
		app_middleware.NewAuthnMiddleware(noAuth),
	)
}

func getRequiredAuth(e *echo.Echo, path string, controller controller) {
	e.GET(path, wrapController(controller),
		app_middleware.NewAuthnMiddleware(requiredAuth),
	)
}

func put(e *echo.Echo, path string, controller controller) {
	e.PUT(path, wrapController(controller),
		app_middleware.NewAuthnMiddleware(requiredAuth),
	)
}

func post(e *echo.Echo, path string, controller controller) {
	e.POST(path, wrapController(controller),
		app_middleware.NewAuthnMiddleware(requiredAuth),
	)
}

func delete(e *echo.Echo, path string, controller controller) {
	e.DELETE(path, wrapController(controller),
		app_middleware.NewAuthnMiddleware(requiredAuth),
	)
}

func wrapController(controller controller) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(context.Context)
		return controller(ctx)
	}
}

func inject(ctx context.Context, req interface{}) error {
	if err := ctx.Bind(req); err != nil {
		return err
	}
	if err := ctx.Validate(req); err != nil {
		return err
	}
	return nil
}

func noAuth(echo.Context) bool {
	return true
}

func requiredAuth(echo.Context) bool {
	return false
}
