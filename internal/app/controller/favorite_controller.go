package controller

import (
	"net/http"

	"github.com/holocycle/holo-back/internal/app/config"
	"github.com/holocycle/holo-back/pkg/api"
	app_context "github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/holocycle/holo-back/pkg/repository"
	"github.com/labstack/echo/v4"
)

type FavoriteController struct {
	Config              *config.AppConfig
	RepositoryContainer *repository.Container
}

func NewFavoriteController(config *config.AppConfig) *FavoriteController {
	return &FavoriteController{
		Config:              config,
		RepositoryContainer: repository.NewContainer(),
	}
}

func (c *FavoriteController) Register(e *echo.Echo) {
	put(e, "/clips/:clip_id/favorite", c.PutFavorite)
	delete(e, "/clips/:clip_id/favorite", c.DeleteFavorite)
}

func (c *FavoriteController) PutFavorite(ctx echo.Context) error {
	goCtx := ctx.Request().Context()

	clipID := ctx.Param("clip_id")
	if clipID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify clip_id")
	}

	if _, err := c.RepositoryContainer.ClipRepository.NewQuery(goCtx).
		Where(&model.Clip{ID: clipID, Status: model.ClipStatusPublic}).
		Find(); err != nil {
		if repository.NotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "clip was not found")
		}
		return err
	}

	favorite := model.NewFavorite(clipID, app_context.GetSession(goCtx).UserID)
	_, err := c.RepositoryContainer.FavoriteRepository.NewQuery(goCtx).Where(favorite).Find()
	if err != nil && !repository.NotFoundError(err) {
		return err
	}
	if err == nil {
		return ctx.JSON(http.StatusConflict, &api.PutFavoriteResponse{})
	}

	if err := c.RepositoryContainer.FavoriteRepository.NewQuery(goCtx).Create(favorite); err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, &api.PutFavoriteResponse{})
}

func (c *FavoriteController) DeleteFavorite(ctx echo.Context) error {
	goCtx := ctx.Request().Context()

	clipID := ctx.Param("clip_id")
	if clipID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify clip_id")
	}

	if _, err := c.RepositoryContainer.ClipRepository.NewQuery(goCtx).
		Where(&model.Clip{ID: clipID, Status: model.ClipStatusPublic}).
		Find(); err != nil {
		if repository.NotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "clip was not found")
		}
		return err
	}

	favorite := model.NewFavorite(clipID, app_context.GetSession(goCtx).UserID)
	rows, err := c.RepositoryContainer.FavoriteRepository.NewQuery(goCtx).Where(favorite).Delete()
	if err != nil {
		return err
	}
	if rows == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "you do not favorite this clip")
	}

	return ctx.JSON(http.StatusOK, &api.DeleteFavoriteRequest{})
}
