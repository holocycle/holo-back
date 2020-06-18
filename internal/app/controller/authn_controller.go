package controller

import (
	"errors"
	"net/http"
	net_url "net/url"
	"time"

	"github.com/holocycle/holo-back/internal/app/config"
	app_context "github.com/holocycle/holo-back/internal/app/context"
	"github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/httpclient"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/holocycle/holo-back/pkg/repository"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type AuthnController struct {
	Config            *config.AppConfig
	UserRepository    repository.UserRepository
	SessionRepository repository.SessionRepository
}

func NewAuthnController(config *config.AppConfig) *AuthnController {
	return &AuthnController{
		Config:            config,
		UserRepository:    repository.NewUserRepository(),
		SessionRepository: repository.NewSessionRepository(),
	}
}

func (c *AuthnController) Register(e *echo.Echo) {
	get(e, "/login/google", c.LoginGoogle)
	get(e, "/login/google-callback", c.LoginGoogleCallback)
	post(e, "/logout", c.Logout)
}

func (c *AuthnController) LoginGoogle(ctx context.Context) error {
	log := ctx.GetLog()

	callbackURL := ctx.FormValue("callback")
	if _, err := net_url.Parse(callbackURL); err != nil || callbackURL == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify parameter `callback`")
	}
	log.Info("paramater is OK", zap.String("callback", callbackURL))

	url, err := httpclient.BuildURL(c.Config.GoogleOAuth2.GoogleAuthURL, map[string]string{
		"client_id":     c.Config.GoogleOAuth2.ClientID,
		"redirect_uri":  c.Config.GoogleOAuth2.ClientRedirectURL,
		"response_type": "code",
		"scope":         c.Config.GoogleOAuth2.Scope,
		"state":         callbackURL,
	})
	if err != nil {
		return err
	}

	return ctx.Redirect(http.StatusFound, url.String())
}

func (c *AuthnController) LoginGoogleCallback(ctx context.Context) error {
	log := ctx.GetLog()

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

	tokenJSON, err := httpclient.Post(c.Config.GoogleOAuth2.GoogleTokenURL, map[string]string{
		"code":          code,
		"client_id":     c.Config.GoogleOAuth2.ClientID,
		"client_secret": c.Config.GoogleOAuth2.ClientSecret,
		"redirect_uri":  c.Config.GoogleOAuth2.ClientRedirectURL,
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

	tokenInfoJSON, err := httpclient.Get(c.Config.GoogleOAuth2.GoogleTokenInfoURL, map[string]string{
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

	tx := ctx.GetDB()
	user, err := c.UserRepository.NewQuery(tx).
		Where(&model.User{Email: email}).Find()
	if err != nil && !repository.NotFoundError(err) {
		return err
	}

	if repository.NotFoundError(err) {
		log.Info("user not found", zap.String("email", email))
		//TODO: imageURLの取得処理
		user = model.NewUser(email, email, "https://yt3.ggpht.com/a/AATXAJwHPp_TkvcWJyblt9XVYDjNSjrj6KdpQSCQNQ=s288-c-k-c0xffffffff-no-rj-mo")
		if err := c.UserRepository.NewQuery(tx).Create(user); err != nil {
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

	if err := c.SessionRepository.NewQuery(tx).Create(session); err != nil {
		return err
	}
	log.Info("success to craete session", zap.Any("session", session))

	callbackURL.Fragment = "token=" + session.ID
	return ctx.Redirect(http.StatusFound, callbackURL.String())
}

func (c *AuthnController) Logout(ctx context.Context) error {
	session := app_context.GetSession(ctx)

	tx := ctx.GetDB()
	_, err := c.SessionRepository.NewQuery(tx).
		Where(&model.Session{ID: session.ID}).Delete()
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}
