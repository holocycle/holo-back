package controller

import "github.com/labstack/echo/v4"

func RegisterAllController(e *echo.Echo) {
	RegisterAppController(e)
	RegisterLiverController(e)
	RegisterAuthnController(e)
	RegisterClipController(e)
}
