package context

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Context interface {
	echo.Context

	SetDB(db *gorm.DB)
	GetDB() *gorm.DB

	SetLog(log *zap.Logger)
	GetLog() *zap.Logger
}

func NewContextFrom(ctx echo.Context) Context {
	return &contextImpl{ctx}
}
