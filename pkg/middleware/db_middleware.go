package middleware

import (
	"fmt"

	"github.com/holocycle/holo-back/pkg/context"
	"go.uber.org/zap"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type DBMiddleware struct {
	DB *gorm.DB
}

func NewDBMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	m := &DBMiddleware{
		DB: db,
	}
	return m.Process
}

func (m *DBMiddleware) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, ok := c.(context.Context)
		if !ok {
			panic("DBMiddleware: require ContextMiddleware before DBMiddleware")
		}

		tx := m.DB.Begin()
		if err := tx.Error; err != nil {
			return err
		}
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
				panic(r)
			}
		}()

		if ctx.GetLog() != nil {
			log := ctx.GetLog()
			tx.SetLogger(&loggerForGorm{log: log})
		}

		ctx.SetDB(tx)
		err := next(ctx)
		if err != nil {
			tx.Rollback()
			if tx.Error != nil {
				fmt.Println("rollback failed", tx.Error) // FIXME: use logger
			}
			return err
		}

		return tx.Commit().Error
	}
}

type loggerForGorm struct {
	log *zap.Logger
}

func (l *loggerForGorm) Print(v ...interface{}) {
	l.log.Debug("gorm log", zap.Any("payload", v))
}
