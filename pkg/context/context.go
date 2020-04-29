package context

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type Context interface {
	echo.Context

	SetDB(db *gorm.DB)
	GetDB() *gorm.DB
}

func NewContextFrom(ctx echo.Context) Context {
	return &contextImpl{ctx}
}
