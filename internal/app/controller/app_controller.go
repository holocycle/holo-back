package controller

import (
	"net/http"

	"github.com/holocycle/holo-back/internal/app/config"
	"github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/model"

	"github.com/labstack/echo/v4"
)

type AppController struct {
	Config *config.AppConfig
}

func NewAppController(config *config.AppConfig) *AppController {
	return &AppController{
		Config: config,
	}
}

func (c *AppController) Register(e *echo.Echo) {
	get(e, "/", c.Index)
	get(e, "/health", c.Health)
}

func (c *AppController) Index(ctx context.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "Hello World",
	})
}

func (c *AppController) Health(ctx context.Context) error {
	tx := ctx.GetDB()

	healthCheck := model.NewHealthCheck()
	if err := tx.Save(healthCheck).Error; err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "OK",
	})
}
