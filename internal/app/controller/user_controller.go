package controller

import (
	"net/http"

	"github.com/holocycle/holo-back/internal/app/config"
	app_context "github.com/holocycle/holo-back/internal/app/context"
	"github.com/holocycle/holo-back/pkg/api"
	"github.com/holocycle/holo-back/pkg/context"
	app_context2 "github.com/holocycle/holo-back/pkg/context2"
	"github.com/holocycle/holo-back/pkg/converter"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/holocycle/holo-back/pkg/repository"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type UserController struct {
	Config             *config.AppConfig
	UserRepository     repository.UserRepository
	FavoriteRepository repository.FavoriteRepository
}

func NewUserController(config *config.AppConfig) *UserController {
	return &UserController{
		Config:             config,
		UserRepository:     repository.NewUserRepository(),
		FavoriteRepository: repository.NewFavoriteRepository(),
	}
}

func (c *UserController) Register(e *echo.Echo) {
	get(e, "/users", c.ListUsers)
	getRequiredAuth(e, "/users/me", c.GetUsersMe)
	getRequiredAuth(e, "/users/me/favorites", c.GetLoginUserFavorites)
	get(e, "/users/:user_id", c.GetOneUser)
	get(e, "/users/:user_id/favorites", c.GetOneUsersFavorites)
}

func (c *UserController) ListUsers(ctx context.Context) error {
	log := ctx.GetLog()
	goCtx := app_context2.FromEchoContext(ctx)

	req := &api.ListUserRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	log.Info("success to validate", zap.Any("req", req))

	query := c.UserRepository.NewQuery(goCtx)
	if req.Limit > 0 {
		query = query.Limit(req.Limit)
	}
	if req.OrderBy == "latest" {
		query = query.Latest()
	}
	users, err := query.FindAll()
	if err != nil {
		return err
	}
	log.Info("success to retrieve users")

	res := make([]*api.User, 0)
	for _, user := range users {
		res = append(res, converter.ConvertToUser(user))
	}

	return ctx.JSON(http.StatusOK, &api.ListUserResponse{
		Users: res,
	})
}

func (c *UserController) GetUsersMe(ctx context.Context) error {
	log := ctx.GetLog()
	goCtx := app_context2.FromEchoContext(ctx)

	req := &api.GetLoginUserRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	log.Info("success to validate", zap.Any("req", req))

	loginUserID := app_context.GetSession(ctx).UserID
	loginUser, err := c.UserRepository.NewQuery(goCtx).Where(&model.User{ID: loginUserID}).Find()
	if err != nil {
		return err
	}

	res := converter.ConvertToLoginUser(loginUser)

	return ctx.JSON(http.StatusCreated, &api.GetLoginUserResponse{
		LoginUser: *res,
	})
}

func (c *UserController) GetLoginUserFavorites(ctx context.Context) error {
	loginUserID := app_context.GetSession(ctx).UserID
	log := ctx.GetLog()
	goCtx := app_context2.FromEchoContext(ctx)

	req := &api.GetUserFavoritesRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	log.Info("success to validate", zap.Any("req", req))

	favorites, err := c.FavoriteRepository.NewQuery(goCtx).
		Where(&model.Favorite{UserID: loginUserID}).
		JoinClip().
		FindAll()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, &api.GetUserFavoritesResponse{
		FavoriteClips: converter.ConvertToFavoriteClips(favorites),
	})
}

func (c *UserController) GetOneUser(ctx context.Context) error {
	log := ctx.GetLog()
	goCtx := app_context2.FromEchoContext(ctx)

	userID := ctx.Param("user_id")
	if userID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify user_id")
	}
	log.Debug("success to validate request", zap.String("userID", userID))

	user, err := c.UserRepository.NewQuery(goCtx).Where(&model.User{ID: userID}).Find()
	if err != nil {
		return err
	}

	res := converter.ConvertToUser(user)

	return ctx.JSON(http.StatusCreated, &api.GetUserResponse{
		User: *res,
	})
}

func (c *UserController) GetOneUsersFavorites(ctx context.Context) error {
	log := ctx.GetLog()
	goCtx := app_context2.FromEchoContext(ctx)

	userID := ctx.Param("user_id")
	if userID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify user_id")
	}
	log.Debug("success to validate request", zap.String("userID", userID))

	favorites, err := c.FavoriteRepository.NewQuery(goCtx).
		Where(&model.Favorite{UserID: userID}).
		JoinClip().
		FindAll()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, &api.GetUserFavoritesResponse{
		FavoriteClips: converter.ConvertToFavoriteClips(favorites),
	})
}
