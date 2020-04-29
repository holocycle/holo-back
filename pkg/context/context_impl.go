package context

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type contextImpl struct {
	echo.Context
}

const (
	keyOfDB = "DATABASE_CONNECTION"
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
