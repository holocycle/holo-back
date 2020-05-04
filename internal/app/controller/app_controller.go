package controller

import (
	"net/http"

	"github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/model"

	"github.com/labstack/echo/v4"
)

func RegisterAppController(e *echo.Echo) {
	e.GET("/", Index)
	e.GET("/health", Health)
}

func Index(c echo.Context) error {
	ctx := c.(context.Context)
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "Hello World",
	})
}

func Health(c echo.Context) error {
	ctx := c.(context.Context)

	db := ctx.GetDB()

	healthCheck := model.NewHealthCheck()
	err := db.Save(healthCheck).Error
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, healthCheck)
}
