package controller

import (
	"net/http"

	"github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/model"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	ctx := c.(context.Context)
	return ctx.String(http.StatusOK, "Index")
}

func Health(c echo.Context) error {
	ctx := c.(context.Context)

	db := ctx.GetDB()

	healthCheck := model.NewHealthCheck()
	err := db.Save(healthCheck).Error
	if err != nil {
		return err
	}

	res := "Health:" + healthCheck.CreatedAt.String()
	return ctx.String(http.StatusOK, res)
}
