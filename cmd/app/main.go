package main

import (
	"fmt"
	"os"

	"github.com/holocycle/holo-back/internal/app/config"
	"github.com/holocycle/holo-back/internal/app/controller"
	"github.com/holocycle/holo-back/pkg/db"
	"github.com/holocycle/holo-back/pkg/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config, err := config.NewConfig()
	if err != nil {
		fmt.Printf("Failed to load configuration. err=%+v\n", err)
		os.Exit(1)
	}

	db, err := db.NewDB(config.DatabaseURL)
	if err != nil {
		fmt.Printf("cannnot access db %+v\n", err) // TODO: use logger
		return
	}

	middlewares := []echo.MiddlewareFunc{
		middleware.NewContextMiddleware(),
		middleware.NewDBMiddleware(db),
	}
	e.Use(middlewares...)

	e.GET("/", controller.Index)
	e.GET("/health", controller.Health)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.Port)))
}
