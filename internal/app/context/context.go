package context

import (
	"github.com/holocycle/holo-back/internal/app/config"
	"github.com/labstack/echo/v4"
)

const (
	appContextConfigKey = "APP_CONTEXT_CONFIG_KEY"
	appContextUserIDKey = "APP_CONTEXT_USER_ID_KEY"
)

func SetConfig(ctx echo.Context, config *config.AppConfig) {
	ctx.Set(appContextConfigKey, config)
}

func GetConfig(ctx echo.Context) *config.AppConfig {
	return ctx.Get(appContextConfigKey).(*config.AppConfig)
}

func SetUserID(ctx echo.Context, userID string) {
	ctx.Set(appContextUserIDKey, userID)
}

func GetUserID(ctx echo.Context) string {
	userID, ok := ctx.Get(appContextUserIDKey).(string)
	if !ok || userID == "" {
		panic("context.GetUserID: context has no userID but called")
	}
	return userID
}
