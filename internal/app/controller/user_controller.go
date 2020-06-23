package controller

import (
	"net/http"

	"github.com/holocycle/holo-back/internal/app/config"
	"github.com/holocycle/holo-back/pkg/api"
	app_context "github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/converter"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/holocycle/holo-back/pkg/repository"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type UserController struct {
	Config              *config.AppConfig
	RepositoryContainer *repository.Container
}

func NewUserController(config *config.AppConfig) *UserController {
	return &UserController{
		Config:              config,
		RepositoryContainer: repository.NewContainer(),
	}
}

func (c *UserController) Register(e *echo.Echo) {
	get(e, "/users", c.ListUsers)
	getRequiredAuth(e, "/users/me", c.GetUsersMe)
	getRequiredAuth(e, "/users/me/favorites", c.GetLoginUserFavorites)
	get(e, "/users/:user_id", c.GetOneUser)
	get(e, "/users/:user_id/favorites", c.GetOneUsersFavorites)
}

func (c *UserController) ListUsers(ctx echo.Context) error {
	goCtx := ctx.Request().Context()
	log := app_context.GetLog(goCtx)

	req := &api.ListUserRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	log.Info("success to validate", zap.Any("req", req))

	query := c.RepositoryContainer.UserRepository.NewQuery(goCtx)
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

func (c *UserController) GetUsersMe(ctx echo.Context) error {
	goCtx := ctx.Request().Context()
	log := app_context.GetLog(goCtx)

	req := &api.GetLoginUserRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	log.Info("success to validate", zap.Any("req", req))

	loginUserID := app_context.GetSession(goCtx).UserID
	loginUser, err := c.RepositoryContainer.UserRepository.NewQuery(goCtx).Where(&model.User{ID: loginUserID}).Find()
	if err != nil {
		return err
	}

	res := converter.ConvertToLoginUser(loginUser)

	return ctx.JSON(http.StatusCreated, &api.GetLoginUserResponse{
		LoginUser: *res,
	})
}

func (c *UserController) GetLoginUserFavorites(ctx echo.Context) error {
	goCtx := ctx.Request().Context()
	log := app_context.GetLog(goCtx)
	loginUserID := app_context.GetSession(goCtx).UserID

	req := &api.GetUserFavoritesRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	log.Info("success to validate", zap.Any("req", req))

	favorites, err := c.RepositoryContainer.FavoriteRepository.NewQuery(goCtx).
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

func (c *UserController) GetOneUser(ctx echo.Context) error {
	goCtx := ctx.Request().Context()
	log := app_context.GetLog(goCtx)

	userID := ctx.Param("user_id")
	if userID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify user_id")
	}
	log.Debug("success to validate request", zap.String("userID", userID))

	user, err := c.RepositoryContainer.UserRepository.NewQuery(goCtx).Where(&model.User{ID: userID}).Find()
	if err != nil {
		return err
	}

	res := converter.ConvertToUser(user)

	return ctx.JSON(http.StatusCreated, &api.GetUserResponse{
		User: *res,
	})
}

func (c *UserController) GetOneUsersFavorites(ctx echo.Context) error {
	goCtx := ctx.Request().Context()
	log := app_context.GetLog(goCtx)

	userID := ctx.Param("user_id")
	if userID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify user_id")
	}
	log.Debug("success to validate request", zap.String("userID", userID))

	favorites, err := c.RepositoryContainer.FavoriteRepository.NewQuery(goCtx).
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
