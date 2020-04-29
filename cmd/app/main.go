package main

import (
	"fmt"

	"github.com/holocycle/holo-back/internal/controller"
	"github.com/holocycle/holo-back/pkg/db"
	"github.com/holocycle/holo-back/pkg/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	dbConf := &db.DBConfig{
		Host:     "db",
		Port:     "5432",
		User:     "holo",
		Password: "password",
	}
	db, err := db.NewDB(dbConf)
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

	e.Logger.Fatal(e.Start(":8080"))
}
