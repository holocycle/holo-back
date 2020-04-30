package controller

import (
	"net/http"

	"github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/holocycle/holo-back/pkg/repository"
	"github.com/labstack/echo/v4"
)

func RegisterLiverController(e *echo.Echo) {
	e.GET("/liver", ListLivers)
	e.GET("/liver/:liverID", GetLiver)
}

func ListLivers(c echo.Context) error {
	ctx := c.(context.Context)

	liverRepo := repository.NewLiverRepository(ctx) // FIXME
	livers, err := liverRepo.FindAll(&model.Liver{})
	if err != nil {
		return err // FIXME
	}

	return ctx.JSON(http.StatusOK, livers)
}

func GetLiver(c echo.Context) error {
	ctx := c.(context.Context)

	liverID := ctx.Param("liverID")

	liverRepo := repository.NewLiverRepository(ctx) // FIXME
	liver, err := liverRepo.FindBy(&model.Liver{ID: liverID})
	if err != nil {
		return err // FIXME
	}

	return ctx.JSON(http.StatusOK, liver)
}
