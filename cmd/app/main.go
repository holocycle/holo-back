package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/holocycle/holo-back/internal/app/config"
	app_context "github.com/holocycle/holo-back/internal/app/context"
	"github.com/holocycle/holo-back/internal/app/controller"
	app_middleware "github.com/holocycle/holo-back/internal/app/middleware"
	"github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/db"
	"github.com/holocycle/holo-back/pkg/logger"
	"github.com/holocycle/holo-back/pkg/middleware"
	"github.com/holocycle/holo-back/pkg/validator"
	"github.com/labstack/echo/v4"
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
	defer log.Sync()
	log.Info("Created logger")

	db, err := db.NewDB(&config.DB)
	if err != nil {
		log.Fatal("cannnot access database", zap.Error(err))
	}
	defer db.Close()
	log.Info("Connected database")

	e.Validator = validator.NewValidator()

	middlewares := []echo.MiddlewareFunc{
		middleware.NewContextMiddleware(),
		middleware.NewContextHandleMiddleware(func(ctx context.Context) (context.Context, error) {
			app_context.SetConfig(ctx, config)
			return ctx, nil
		}),
		middleware.NewLoggerMiddleware(log),
		middleware.NewRequestLoggingMiddleware(),
		middleware.NewErrorLoggingMiddleware(),
		middleware.NewResponseLoggingMiddleware(),
		middleware.NewDBMiddleware(db),
		app_middleware.NewAuthnMiddleware(func(ctx echo.Context) bool {
			return ctx.Request().Method == http.MethodGet
		}),
	}
	e.Use(middlewares...)

	e.Static("/assets", "assets")
	controller.RegisterController(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.Port)))
}
