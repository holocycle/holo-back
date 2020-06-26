package controller

import (
	"net/http"
	"strconv"

	"github.com/holocycle/holo-back/internal/app/config"
	"github.com/holocycle/holo-back/pkg/api"
	"github.com/holocycle/holo-back/pkg/service"
	"github.com/labstack/echo/v4"
)

type CliplistController struct {
	Config           *config.AppConfig
	ServiceContainer *service.Container
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

func (c *CliplistController) ListCliplists(ctx echo.Context) error {
	req := &api.ListCliplistsRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := c.ServiceContainer.CliplistService.ListCliplists(
		ctx.Request().Context(),
		req,
	)
	if err != nil {
		return echo.NewHTTPError(ConvertToStatus(err), err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *CliplistController) GetCliplist(ctx echo.Context) error {
	req := &api.GetCliplistRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	cliplistID := ctx.Param("cliplist_id")
	if cliplistID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify cliplist_id")
	}

	res, err := c.ServiceContainer.CliplistService.GetCliplist(
		ctx.Request().Context(),
		cliplistID,
		req,
	)
	if err != nil {
		return echo.NewHTTPError(ConvertToStatus(err), err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *CliplistController) PostCliplist(ctx echo.Context) error {
	req := &api.PostCliplistRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := c.ServiceContainer.CliplistService.PostCliplist(
		ctx.Request().Context(),
		req,
	)
	if err != nil {
		return echo.NewHTTPError(ConvertToStatus(err), err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *CliplistController) PutCliplist(ctx echo.Context) error {
	req := &api.PutCliplistRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	cliplistID := ctx.Param("cliplist_id")
	if cliplistID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify cliplist_id")
	}

	res, err := c.ServiceContainer.CliplistService.PutCliplist(
		ctx.Request().Context(),
		cliplistID,
		req,
	)
	if err != nil {
		return echo.NewHTTPError(ConvertToStatus(err), err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *CliplistController) DeleteCliplist(ctx echo.Context) error {
	req := &api.DeleteCliplistRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	cliplistID := ctx.Param("cliplist_id")
	if cliplistID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify cliplist_id")
	}

	res, err := c.ServiceContainer.CliplistService.DeleteCliplist(
		ctx.Request().Context(),
		cliplistID,
		req,
	)
	if err != nil {
		return echo.NewHTTPError(ConvertToStatus(err), err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *CliplistController) GetCliplistItem(ctx echo.Context) error {
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
		ctx.Request().Context(),
		cliplistID,
		index,
		req,
	)
	if err != nil {
		return echo.NewHTTPError(ConvertToStatus(serr), err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *CliplistController) PostCliplistItem(ctx echo.Context) error {
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
		ctx.Request().Context(),
		cliplistID,
		index,
		req,
	)
	if err != nil {
		return echo.NewHTTPError(ConvertToStatus(serr), err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *CliplistController) DeleteCliplistItem(ctx echo.Context) error {
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
		ctx.Request().Context(),
		cliplistID,
		index,
		req,
	)
	if err != nil {
		return echo.NewHTTPError(ConvertToStatus(serr), err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}
