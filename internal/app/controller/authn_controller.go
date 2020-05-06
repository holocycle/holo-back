package controller

import (
	"errors"
	"net/http"
	net_url "net/url"
	"time"

	app_context "github.com/holocycle/holo-back/internal/app/context"
	"github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/httpclient"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/holocycle/holo-back/pkg/repository"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func RegisterAuthnController(e *echo.Echo) {
	e.GET("/login/google", LoginGoogle)
	e.GET("/login/google-callback", LoginGoogleCallback)
	e.POST("/logout", Logout)
}

func LoginGoogle(c echo.Context) error {
	ctx := c.(context.Context)
	cfg := app_context.GetConfig(ctx)
	log := ctx.GetLog()

	callbackURL := ctx.FormValue("callback")
	if _, err := net_url.Parse(callbackURL); err != nil || callbackURL == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify parameter `callback`")
	}
	log.Info("paramater is OK", zap.String("callback", callbackURL))

	url, err := httpclient.BuildURL(cfg.GoogleOAuth2.GoogleAuthURL, map[string]string{
		"client_id":     cfg.GoogleOAuth2.ClientID,
		"redirect_uri":  cfg.GoogleOAuth2.ClientRedirectURL,
		"response_type": "code",
		"scope":         cfg.GoogleOAuth2.Scope,
		"state":         callbackURL,
	})
	if err != nil {
		return err
	}

	return ctx.Redirect(http.StatusFound, url.String())
}

func LoginGoogleCallback(c echo.Context) error {
	ctx := c.(context.Context)
	log := ctx.GetLog()
	cfg := app_context.GetConfig(ctx)

	code := ctx.FormValue("code")
	if code == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify parameter `code`")
	}

	callbackURLText := ctx.FormValue("state")
	callbackURL, err := net_url.Parse(callbackURLText)
	if err != nil || callbackURLText == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify parameter `state`")
	}
	log.Info("paramater is OK",
		zap.String("code", code),
		zap.String("callbackURL", callbackURL.String()))

	tokenJSON, err := httpclient.Post(cfg.GoogleOAuth2.GoogleTokenURL, map[string]string{
		"code":          code,
		"client_id":     cfg.GoogleOAuth2.ClientID,
		"client_secret": cfg.GoogleOAuth2.ClientSecret,
		"redirect_uri":  cfg.GoogleOAuth2.ClientRedirectURL,
		"grant_type":    "authorization_code",
	})
	if err != nil {
		return err
	}

	idToken, ok := tokenJSON["id_token"].(string)
	if !ok {
		log.Error("id_token not found", zap.Any("response", tokenJSON))
		return errors.New("id token not found")
	}
	log.Info("success to retrieve token", zap.String("idToken", idToken))

	tokenInfoJSON, err := httpclient.Get(cfg.GoogleOAuth2.GoogleTokenInfoURL, map[string]string{
		"id_token": idToken,
	})
	if err != nil {
		return err
	}

	email, ok := tokenInfoJSON["email"].(string)
	if !ok {
		log.Error("email not found", zap.Any("response", tokenInfoJSON))
		return errors.New("email not found")
	}
	log.Info("success to retrieve email info", zap.String("email", email))

	userRepo := repository.NewUserRepository(ctx)
	user, err := userRepo.FindBy(&model.User{Email: email})
	if err != nil && !repository.NotFoundError(err) {
		return err
	}

	if repository.NotFoundError(err) {
		log.Info("user not found", zap.String("email", email))
		user = model.NewUser(email, email)
		if err := userRepo.Create(user); err != nil {
			return err
		}
		log.Info("success to create user", zap.Any("user", user))
	}
	log.Info("success to find user", zap.Any("user", user))

	tokenDuration, err := time.ParseDuration("600s") // FIXME
	if err != nil {
		return err // FIXME
	}
	expireAt := time.Now().Add(tokenDuration)
	session := model.NewSession(user.ID, &expireAt)

	sessionRepo := repository.NewSessionRepository(ctx)
	if err := sessionRepo.Create(session); err != nil {
		return err
	}
	log.Info("success to craete session", zap.Any("session", session))

	callbackURL.Fragment = "token=" + session.ID
	return ctx.Redirect(http.StatusFound, callbackURL.String())
}

func Logout(c echo.Context) error {
	ctx := c.(context.Context)

	session := app_context.GetSession(ctx)

	sessionRepo := repository.NewSessionRepository(ctx)
	if err := sessionRepo.Delete(session); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}
