package controller

import (
	"net/http"
	"strconv"

	"github.com/holocycle/holo-back/internal/app/config"
	"github.com/holocycle/holo-back/pkg/api"
	"github.com/holocycle/holo-back/pkg/context"
	app_context2 "github.com/holocycle/holo-back/pkg/context2"
	"github.com/holocycle/holo-back/pkg/repository"
	"github.com/holocycle/holo-back/pkg/service"
	"github.com/labstack/echo/v4"
)

type CliplistController struct {
	Config                    *config.AppConfig
	ClipRepository            repository.ClipRepository
	CliplistRepository        repository.CliplistRepository
	CliplistContainRepository repository.CliplistContainRepository
	ServiceContainer          *service.Container
}

func NewCliplistController(config *config.AppConfig) *CliplistController {
	return &CliplistController{
		Config:           config,
		ServiceContainer: service.NewContainer(),
	}
}

func (c *CliplistController) Register(e *echo.Echo) {
	get(e, "/cliplists", c.ListCliplists)
	get(e, "/cliplists/:cliplist_id", c.GetCliplist)
	post(e, "/cliplists", c.PostCliplist)
	put(e, "/cliplists/:cliplist_id", c.PutCliplist)
	delete(e, "/cliplists/:cliplist_id", c.DeleteCliplist)

	get(e, "/cliplists/:cliplist_id/:index", c.GetCliplistItem)
	put(e, "/cliplists/:cliplist_id/:index", c.PostCliplistItem)
	delete(e, "/cliplists/:cliplist_id/:index", c.DeleteCliplistItem)
}

func (c *CliplistController) ListCliplists(ctx context.Context) error {
	req := &api.ListCliplistsRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := c.ServiceContainer.CliplistService.ListCliplists(
		app_context2.FromEchoContext(ctx),
		req,
	)
	if err != nil {
		return echo.NewHTTPError(ConvertToStatus(err), err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *CliplistController) GetCliplist(ctx context.Context) error {
	req := &api.GetCliplistRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	cliplistID := ctx.Param("cliplist_id")
	if cliplistID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify cliplist_id")
	}

	res, err := c.ServiceContainer.CliplistService.GetCliplist(
		app_context2.FromEchoContext(ctx),
		cliplistID,
		req,
	)
	if err != nil {
		return echo.NewHTTPError(ConvertToStatus(err), err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *CliplistController) PostCliplist(ctx context.Context) error {
	req := &api.PostCliplistRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := c.ServiceContainer.CliplistService.PostCliplist(
		app_context2.FromEchoContext(ctx),
		req,
	)
	if err != nil {
		return echo.NewHTTPError(ConvertToStatus(err), err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *CliplistController) PutCliplist(ctx context.Context) error {
	req := &api.PutCliplistRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	cliplistID := ctx.Param("cliplist_id")
	if cliplistID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify cliplist_id")
	}

	res, err := c.ServiceContainer.CliplistService.PutCliplist(
		app_context2.FromEchoContext(ctx),
		cliplistID,
		req,
	)
	if err != nil {
		return echo.NewHTTPError(ConvertToStatus(err), err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *CliplistController) DeleteCliplist(ctx context.Context) error {
	req := &api.DeleteCliplistRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	cliplistID := ctx.Param("cliplist_id")
	if cliplistID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify cliplist_id")
	}

	res, err := c.ServiceContainer.CliplistService.DeleteCliplist(
		app_context2.FromEchoContext(ctx),
		cliplistID,
		req,
	)
	if err != nil {
		return echo.NewHTTPError(ConvertToStatus(err), err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *CliplistController) GetCliplistItem(ctx context.Context) error {
	req := &api.GetCliplistItemRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	cliplistID := ctx.Param("cliplist_id")
	if cliplistID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify cliplist_id")
	}

	indexInStr := ctx.Param("index")
	if indexInStr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify index")
	}
	index, err := strconv.Atoi(indexInStr)
	if err != nil || index < 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "index must be zero or positive number")
	}

	res, serr := c.ServiceContainer.CliplistItemService.GetCliplistItem(
		app_context2.FromEchoContext(ctx),
		cliplistID,
		index,
		req,
	)
	if err != nil {
		return echo.NewHTTPError(ConvertToStatus(serr), err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *CliplistController) PostCliplistItem(ctx context.Context) error {
	req := &api.PostCliplistItemRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	cliplistID := ctx.Param("cliplist_id")
	if cliplistID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify cliplist_id")
	}

	indexInStr := ctx.Param("index")
	if indexInStr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify index")
	}
	index, err := strconv.Atoi(indexInStr)
	if err != nil || index < 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "index must be zero or positive number")
	}

	res, serr := c.ServiceContainer.CliplistItemService.PostCliplistItem(
		app_context2.FromEchoContext(ctx),
		cliplistID,
		index,
		req,
	)
	if err != nil {
		return echo.NewHTTPError(ConvertToStatus(serr), err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *CliplistController) DeleteCliplistItem(ctx context.Context) error {
	req := &api.DeleteCliplistItemRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	cliplistID := ctx.Param("cliplist_id")
	if cliplistID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify cliplist_id")
	}

	indexInStr := ctx.Param("index")
	if indexInStr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify index")
	}
	index, err := strconv.Atoi(indexInStr)
	if err != nil || index < 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "index must be zero or positive number")
	}

	res, serr := c.ServiceContainer.CliplistItemService.DeleteCliplistItem(
		app_context2.FromEchoContext(ctx),
		cliplistID,
		index,
		req,
	)
	if err != nil {
		return echo.NewHTTPError(ConvertToStatus(serr), err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}
