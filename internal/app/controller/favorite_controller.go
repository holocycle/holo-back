package controller

import (
	"net/http"

	"github.com/holocycle/holo-back/internal/app/config"
	app_context "github.com/holocycle/holo-back/internal/app/context"
	"github.com/holocycle/holo-back/pkg/api"
	"github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/holocycle/holo-back/pkg/repository"
	"github.com/labstack/echo/v4"
)

type FavoriteController struct {
	Config             *config.AppConfig
	ClipRepository     repository.ClipRepository
	FavoriteRepository repository.FavoriteRepository
}

func NewFavoriteController(config *config.AppConfig) *FavoriteController {
	return &FavoriteController{
		Config:             config,
		ClipRepository:     repository.NewClipRepository(),
		FavoriteRepository: repository.NewFavoriteRepository(),
	}
}

func (c *FavoriteController) Register(e *echo.Echo) {
	put(e, "/clips/:clip_id/favorite", c.PutFavorite)
	delete(e, "/clips/:clip_id/favorite", c.DeleteFavorite)
}

func (c *FavoriteController) PutFavorite(ctx context.Context) error {
	clipID := ctx.Param("clip_id")
	if clipID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify clip_id")
	}

	if _, err := c.ClipRepository.NewQuery(ctx.GetDB()).
		Where(&model.Clip{ID: clipID}).Find(); err != nil {
		if repository.NotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "clip was not found")
		}
		return err
	}

	favorite := model.NewFavorite(clipID, app_context.GetUserID(ctx))
	_, err := c.FavoriteRepository.NewQuery(ctx.GetDB()).Where(favorite).Find()
	if err != nil && !repository.NotFoundError(err) {
		return err
	}
	if err == nil {
		return ctx.JSON(http.StatusConflict, &api.PutFavoriteResponse{})
	}

	if err := c.FavoriteRepository.NewQuery(ctx.GetDB()).Create(favorite); err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, &api.PutFavoriteResponse{})
}

func (c *FavoriteController) DeleteFavorite(ctx context.Context) error {
	clipID := ctx.Param("clip_id")
	if clipID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify clip_id")
	}

	if _, err := c.ClipRepository.NewQuery(ctx.GetDB()).
		Where(&model.Clip{ID: clipID}).Find(); err != nil {
		if repository.NotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "clip was not found")
		}
		return err
	}

	favorite := model.NewFavorite(clipID, app_context.GetUserID(ctx))
	rows, err := c.FavoriteRepository.NewQuery(ctx.GetDB()).Where(favorite).Delete()
	if err != nil {
		return err
	}
	if rows == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "you do not favorite this clip")
	}

	return ctx.JSON(http.StatusOK, &api.DeleteFavoriteRequest{})
}
