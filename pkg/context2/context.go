package context

import (
	"context"

	echo_context "github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

const (
	keyOfDB      = "DATABASE_CONNECTION"
	keyOfLogger  = "LOGGER"
	keyOfSession = "SESSION"
)

func SetDB(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, keyOfDB, db)
}

func GetDB(ctx context.Context) *gorm.DB {
	val := ctx.Value(keyOfDB)
	if val == nil {
		panic("context.GetDB: no db connection in context.")
	}
	return val.(*gorm.DB)
}

func SetLog(ctx context.Context, log *zap.Logger) context.Context {
	return context.WithValue(ctx, keyOfLogger, log)
}

func GetLog(ctx context.Context) *zap.Logger {
	val := ctx.Value(keyOfLogger)
	if val == nil {
		panic("context.GetLog: no logger in context.")
	}
	return val.(*zap.Logger)
}

func SetSession(ctx context.Context, session *model.Session) context.Context {
	return context.WithValue(ctx, keyOfSession, session)
}

func GetSession(ctx context.Context) *model.Session {
	val := ctx.Value(keyOfSession)
	if val == nil {
		panic("context.GetSession: no session in context.")
	}
	return val.(*model.Session)
}

func FromEchoContext(ctx echo_context.Context) context.Context {
	res := ctx.Request().Context()
	res = SetDB(res, ctx.GetDB())
	res = SetLog(res, ctx.GetLog())
	return res
}
