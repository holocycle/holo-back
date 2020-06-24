package controller

import (
	"net/http"

	"github.com/holocycle/holo-back/internal/app/config"
	app_context "github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/service"
	"github.com/labstack/echo/v4"
)

type FavoriteController struct {
	Config           *config.AppConfig
	ServiceContainer *service.Container
}

func NewFavoriteController(config *config.AppConfig) *FavoriteController {
	return &FavoriteController{
		Config:           config,
		ServiceContainer: service.NewContainer(),
	}
}

func (c *FavoriteController) Register(e *echo.Echo) {
	get(e, "/clips/:clip_id/favorite", c.GetFavorite)
	put(e, "/clips/:clip_id/favorite", c.PutFavorite)
	delete(e, "/clips/:clip_id/favorite", c.DeleteFavorite)
}

func (c *FavoriteController) GetFavorite(ctx echo.Context) error {
	goCtx := ctx.Request().Context()

	clipID := ctx.Param("clip_id")
	if clipID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify clip_id")
	}

	userID := app_context.GetSession(goCtx).UserID
	if userID == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "please login")
	}

	res, err := c.ServiceContainer.FavoriteService.GetFavoriteItem(
		ctx.Request().Context(),
		clipID,
		userID,
	)

	if err != nil {
		return echo.NewHTTPError(ConvertToStatus(err), err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *FavoriteController) PutFavorite(ctx echo.Context) error {
	goCtx := ctx.Request().Context()

	clipID := ctx.Param("clip_id")
	if clipID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify clip_id")
	}

	userID := app_context.GetSession(goCtx).UserID
	if userID == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "please login")
	}

	res, err := c.ServiceContainer.FavoriteService.PutFavoriteItem(
		ctx.Request().Context(),
		clipID,
		userID,
	)

	if err != nil {
		return echo.NewHTTPError(ConvertToStatus(err), err.Error())
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (c *FavoriteController) DeleteFavorite(ctx echo.Context) error {
	goCtx := ctx.Request().Context()

	clipID := ctx.Param("clip_id")
	if clipID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify clip_id")
	}

	userID := app_context.GetSession(goCtx).UserID
	if userID == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "please login")
	}

	res, err := c.ServiceContainer.FavoriteService.DeleteFavoriteItem(
		ctx.Request().Context(),
		clipID,
		userID,
	)

	if err != nil {
		return echo.NewHTTPError(ConvertToStatus(err), err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}
