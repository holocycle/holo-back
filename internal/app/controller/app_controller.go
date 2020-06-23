package controller

import (
	"net/http"

	"github.com/holocycle/holo-back/internal/app/config"
	app_context "github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/holocycle/holo-back/pkg/repository"
	"github.com/pkg/errors"

	"github.com/labstack/echo/v4"
)

type AppController struct {
	Config              *config.AppConfig
	RepositoryContainer *repository.Container
}

func NewAppController(config *config.AppConfig) *AppController {
	return &AppController{
		Config:              config,
		RepositoryContainer: repository.NewContainer(),
	}
}

func (c *AppController) Register(e *echo.Echo) {
	get(e, "/", c.Index)
	get(e, "/health", c.Health)
}

func (c *AppController) Index(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "Hello World",
	})
}

func (c *AppController) Health(ctx echo.Context) error {
	tx := app_context.GetDB(ctx.Request().Context())

	healthCheck := model.NewHealthCheck()
	if err := tx.Save(healthCheck).Error; err != nil {
		return errors.WithStack(err)
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "OK",
	})
}
