package middleware

import (
	"net/http"
	"strings"
	"time"

	app_context "github.com/holocycle/holo-back/internal/app/context"
	"github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/holocycle/holo-back/pkg/repository"

	"github.com/labstack/echo/v4"
	echo_middleware "github.com/labstack/echo/v4/middleware"
)

const (
	headerKey = "Authorization"
)

type AuthnMiddleware struct {
	Skipper echo_middleware.Skipper
}

func NewAuthnMiddleware(skipper echo_middleware.Skipper) echo.MiddlewareFunc {
	m := &AuthnMiddleware{
		Skipper: skipper,
	}
	return m.Process
}

func (m *AuthnMiddleware) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(context.Context)
		if m.Skipper(c) {
			return next(ctx)
		}

		token, err := getAuthToken(ctx.Request().Header)
		if err != nil {
			return err
		}

		tx := ctx.GetDB()
		session, err := repository.NewSessionRepository().NewQuery(tx).
			Where(&model.Session{ID: token}).Find()
		if err != nil {
			if repository.NotFoundError(err) {
				return echo.NewHTTPError(http.StatusUnauthorized, "`Authorization` token is invalid")
			}
			return err
		}

		currentTime := time.Now()
		if currentTime.After(*session.ExpireAt) {
			return echo.NewHTTPError(http.StatusUnauthorized, "`Authorization` token is expired")
		}

		app_context.SetSession(ctx, session)
		return next(ctx)
	}
}

func getAuthToken(header http.Header) (string, error) {
	// header := { "Authorization": ["Bearer XYZ"] }
	// values := ["Bearer XYZ"]
	values, ok := header[headerKey]
	if !ok || len(values) == 0 {
		return "", echo.NewHTTPError(http.StatusUnauthorized, "`Authorization` header is not found")
	}

	// authValue := "Bearer XYZ"
	authValue := values[0]

	// authValues := [ "Bearer", "XYZ" ]
	authValues := strings.Split(authValue, " ")
	if len(authValues) != 2 || authValues[0] != "Bearer" {
		return "", echo.NewHTTPError(http.StatusUnauthorized, "invalid `Authorization` header format")
	}

	// token := "XYZ"
	token := authValues[1]
	return token, nil
}
