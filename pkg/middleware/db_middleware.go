package middleware

import (
	"fmt"

	"go.uber.org/zap"

	app_context "github.com/holocycle/holo-back/pkg/context"
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

		log := app_context.GetLog(c.Request().Context())
		tx.SetLogger(&loggerForGorm{log: log})

		newCtx := app_context.SetDB(c.Request().Context(), tx)
		c.SetRequest(c.Request().WithContext(newCtx))
		err := next(c)
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
