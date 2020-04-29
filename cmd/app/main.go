package main

import (
	"fmt"
	"os"

	"github.com/holocycle/holo-back/internal/controller"
	"github.com/holocycle/holo-back/pkg/db"
	"github.com/holocycle/holo-back/pkg/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	dbURL := os.Getenv("DATABASE_URL")
	db, err := db.NewDB(dbURL)
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

	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + port))
}
