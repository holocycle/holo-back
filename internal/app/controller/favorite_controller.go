package controller

import (
	"net/http"

	"github.com/holocycle/holo-back/internal/app/config"
	"github.com/holocycle/holo-back/pkg/api"
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
	getRequiredAuth(e, "/clips/:clip_id/favorite", c.GetFavorite)
	put(e, "/clips/:clip_id/favorite", c.PutFavorite)
	delete(e, "/clips/:clip_id/favorite", c.DeleteFavorite)
}

func (c *FavoriteController) GetFavorite(ctx echo.Context) error {
	req := &api.GetFavoriteRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	goCtx := ctx.Request().Context()

	clipID := ctx.Param("clip_id")
	if clipID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify clip_id")
	}

	res, err := c.ServiceContainer.FavoriteService.GetFavorite(goCtx, clipID, req)

	if err != nil {
		return echo.NewHTTPError(ConvertToStatus(err), err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *FavoriteController) PutFavorite(ctx echo.Context) error {
	req := &api.PutFavoriteRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	goCtx := ctx.Request().Context()

	clipID := ctx.Param("clip_id")
	if clipID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify clip_id")
	}

	res, err := c.ServiceContainer.FavoriteService.PutFavorite(goCtx, clipID, req)

	if err != nil {
		return echo.NewHTTPError(ConvertToStatus(err), err.Error())
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (c *FavoriteController) DeleteFavorite(ctx echo.Context) error {
	req := &api.DeleteFavoriteRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	goCtx := ctx.Request().Context()

	clipID := ctx.Param("clip_id")
	if clipID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify clip_id")
	}

	res, err := c.ServiceContainer.FavoriteService.DeleteFavorite(goCtx, clipID, req)

	if err != nil {
		return echo.NewHTTPError(ConvertToStatus(err), err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}
