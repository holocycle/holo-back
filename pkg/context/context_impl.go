package context

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type contextImpl struct {
	echo.Context
}

const (
	keyOfDB     = "DATABASE_CONNECTION"
	keyOfLogger = "LOGGER"
)

func (ctx *contextImpl) SetDB(db *gorm.DB) {
	ctx.Set(keyOfDB, db)
}

func (ctx *contextImpl) GetDB() *gorm.DB {
	val := ctx.Get(keyOfDB)
	if val == nil {
		panic("context.GetDB: no db connection in context.")
	}
	return val.(*gorm.DB)
}

func (ctx *contextImpl) SetLog(log *zap.Logger) {
	ctx.Set(keyOfLogger, log)
}

func (ctx *contextImpl) GetLog() *zap.Logger {
	val := ctx.Get(keyOfLogger)
	if val == nil {
		panic("context.GetLog: no logger in context.")
	}
	return val.(*zap.Logger)
}
