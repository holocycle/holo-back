package controller

import (
	app_middleware "github.com/holocycle/holo-back/internal/app/middleware"
	"github.com/labstack/echo/v4"
)

func get(e *echo.Echo, path string, controller echo.HandlerFunc) {
	e.GET(path, controller)
}

func getRequiredAuth(e *echo.Echo, path string, controller echo.HandlerFunc) {
	e.GET(path, controller,
		app_middleware.NewAuthnMiddleware(requiredAuth),
	)
}

func put(e *echo.Echo, path string, controller echo.HandlerFunc) {
	e.PUT(path, controller,
		app_middleware.NewAuthnMiddleware(requiredAuth),
	)
}

func post(e *echo.Echo, path string, controller echo.HandlerFunc) {
	e.POST(path, controller,
		app_middleware.NewAuthnMiddleware(requiredAuth),
	)
}

func delete(e *echo.Echo, path string, controller echo.HandlerFunc) {
	e.DELETE(path, controller,
		app_middleware.NewAuthnMiddleware(requiredAuth),
	)
}

func inject(ctx echo.Context, req interface{}) error {
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
