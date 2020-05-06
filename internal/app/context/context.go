package context

import (
	"github.com/holocycle/holo-back/internal/app/config"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/labstack/echo/v4"
)

const (
	appContextConfigKey  = "APP_CONTEXT_CONFIG_KEY"
	appContextSessionKey = "APP_CONTEXT_SESSION_KEY"
)

func SetConfig(ctx echo.Context, config *config.AppConfig) {
	ctx.Set(appContextConfigKey, config)
}

func GetConfig(ctx echo.Context) *config.AppConfig {
	return ctx.Get(appContextConfigKey).(*config.AppConfig)
}

func SetSession(ctx echo.Context, session *model.Session) {
	ctx.Set(appContextSessionKey, session)
}

func GetSession(ctx echo.Context) *model.Session {
	session, ok := ctx.Get(appContextSessionKey).(*model.Session)
	if !ok || session == nil {
		panic("context.GetSession: context has no session but called")
	}
	return session
}
