package main

import (
	"context"
	"fmt"
	"os"

	"github.com/holocycle/holo-back/internal/app/config"
	"github.com/holocycle/holo-back/internal/app/controller"
	app_context "github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/db"
	"github.com/holocycle/holo-back/pkg/logger"
	"github.com/holocycle/holo-back/pkg/middleware"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/holocycle/holo-back/pkg/validator"

	"github.com/labstack/echo/v4"
	echo_middleware "github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func main() {
	e := echo.New()

	config, err := config.NewConfig()
	if err != nil {
		fmt.Printf("Failed to load configuration. err=%+v\n", err)
		os.Exit(1)
	}

	log, err := logger.NewLogger(&config.Logger)
	if err != nil {
		fmt.Printf("Failed to create logger. err=%+v\n", err)
		os.Exit(1)
	}
	defer func() {
		if log.Sync() != nil {
			log.Error("failed log sync", zap.Error(err))
		}
	}()
	log.Info("Created logger")

	db, err := db.NewDB(&config.DB)
	if err != nil {
		log.Fatal("cannnot access database", zap.Error(err))
	}
	defer db.Close()
	log.Info("Connected database")

	e.Validator = validator.NewValidator()

	middlewares := []echo.MiddlewareFunc{
		echo_middleware.Recover(),
		middleware.NewCORSMiddleware(&config.CORS),
		middleware.NewLoggerMiddleware(log),
		middleware.NewContextHandleMiddleware(func(ctx context.Context) (context.Context, error) {
			id := model.GetIDGenerator().New()
			log := app_context.GetLog(ctx).With(zap.String("requestID", id))
			return app_context.SetLog(ctx, log), nil
		}),
		middleware.NewRequestLoggingMiddleware(),
		middleware.NewErrorLoggingMiddleware(),
		middleware.NewResponseLoggingMiddleware(),
		middleware.NewDBMiddleware(db),
	}
	e.Use(middlewares...)

	e.Static("/assets", "assets")
	controller.NewAppController(config).Register(e)
	controller.NewAuthnController(config).Register(e)
	controller.NewLiverController(config).Register(e)
	controller.NewClipController(config).Register(e)
	controller.NewCommentController(config).Register(e)
	controller.NewTagController(config).Register(e)
	controller.NewUserController(config).Register(e)
	controller.NewFavoriteController(config).Register(e)
	controller.NewCliplistController(config).Register(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.Port)))
}
