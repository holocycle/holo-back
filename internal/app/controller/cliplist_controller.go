package controller

import (
	"net/http"
	"strconv"

	"github.com/holocycle/holo-back/internal/app/config"
	app_context "github.com/holocycle/holo-back/internal/app/context"
	"github.com/holocycle/holo-back/pkg/api"
	"github.com/holocycle/holo-back/pkg/context"
	app_context2 "github.com/holocycle/holo-back/pkg/context2"
	"github.com/holocycle/holo-back/pkg/converter"
	"github.com/holocycle/holo-back/pkg/model"
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
		Config:                    config,
		ClipRepository:            repository.NewClipRepository(),
		CliplistRepository:        repository.NewCliplistRepository(),
		CliplistContainRepository: repository.NewCliplistContainRepository(),
		ServiceContainer:          service.NewContainer(),
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
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := c.ServiceContainer.CliplistService.ListCliplists(
		app_context2.FromEchoContext(ctx),
		req,
	)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *CliplistController) GetCliplist(ctx context.Context) error {
	cliplistID := ctx.Param("cliplist_id")
	if cliplistID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify cliplist_id")
	}

	req := &api.GetCliplistRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	cliplist, err := c.CliplistRepository.NewQuery(ctx.GetDB()).
		JoinClip().
		Where(&model.Cliplist{
			ID:     cliplistID,
			Status: model.CliplistStatusPublic,
		}).
		Find()
	if err != nil {
		if repository.NotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "cliplist was not found")
		}
		return err
	}

	pageBegin := req.ItemPerPage * req.Page
	pageEnd := pageBegin + req.ItemPerPage
	if pageEnd > len(cliplist.CliplistContains) {
		pageEnd = len(cliplist.CliplistContains)
	}

	return ctx.JSON(http.StatusOK, &api.GetCliplistResponse{
		Cliplist:      converter.ConvertToCliplist(cliplist),
		PageInfo:      converter.ConvertToPageInfo(len(cliplist.CliplistContains), req.Page, req.ItemPerPage),
		CliplistItems: converter.ConvertToCliplistItems(cliplist.CliplistContains[pageBegin:pageEnd]),
	})
}

func (c *CliplistController) PostCliplist(ctx context.Context) error {
	req := &api.PostCliplistRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	cliplist := model.NewCliplist(
		app_context.GetUserID(ctx),
		req.Title,
		req.Description,
		model.CliplistStatusPublic,
	)
	err := c.CliplistRepository.NewQuery(ctx.GetDB()).Create(cliplist)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, &api.PostCliplistResponse{
		CliplistID: cliplist.ID,
	})
}

func (c *CliplistController) PutCliplist(ctx context.Context) error {
	cliplistID := ctx.Param("cliplist_id")
	if cliplistID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify cliplist_id")
	}

	req := &api.PostCliplistRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	cliplist, err := c.CliplistRepository.NewQuery(ctx.GetDB()).
		Where(&model.Cliplist{
			ID:     cliplistID,
			Status: model.CliplistStatusPublic,
		}).
		Find()
	if err != nil {
		if repository.NotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "cliplist was not found")
		}
		return err
	}
	if cliplist.UserID != app_context.GetUserID(ctx) {
		return echo.NewHTTPError(http.StatusForbidden, "cliplist is not yours")
	}

	cliplist.Title = req.Title
	cliplist.Description = req.Description
	if err = c.CliplistRepository.NewQuery(ctx.GetDB()).Save(cliplist); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, &api.PostCliplistResponse{
		CliplistID: cliplist.ID,
	})
}

func (c *CliplistController) DeleteCliplist(ctx context.Context) error {
	cliplistID := ctx.Param("cliplist_id")
	if cliplistID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify cliplist_id")
	}

	cliplist, err := c.CliplistRepository.NewQuery(ctx.GetDB()).
		Where(&model.Cliplist{
			ID:     cliplistID,
			Status: model.CliplistStatusPublic,
		}).
		Find()
	if err != nil {
		if repository.NotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "cliplist was not found")
		}
		return err
	}
	if cliplist.UserID != app_context.GetUserID(ctx) {
		return echo.NewHTTPError(http.StatusForbidden, "cliplist is not yours")
	}

	cliplist.Status = model.CliplistStatusPublic
	if err := c.CliplistRepository.NewQuery(ctx.GetDB()).Save(cliplist); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, &api.DeleteCliplistResponse{})
}

func (c *CliplistController) GetCliplistItem(ctx context.Context) error {
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

	_, err = c.CliplistRepository.NewQuery(ctx.GetDB()).
		Where(&model.Cliplist{
			Status: model.CliplistStatusPublic,
		}).
		Find()
	if err != nil {
		if repository.NotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "cliplist was not found")
		}
		return err
	}

	cliplistContain, err := c.CliplistContainRepository.NewQuery(ctx.GetDB()).
		JoinClip().
		Where(&model.CliplistContain{
			CliplistID: cliplistID,
			Index:      index,
		}).Find()
	if err != nil {
		if repository.NotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "cliplist index was out of range")
		}
		return err
	}

	return ctx.JSON(http.StatusOK, &api.GetCliplistItemResponse{
		CliplistItem: converter.ConvertToCliplistItem(cliplistContain),
	})
}

func (c *CliplistController) PostCliplistItem(ctx context.Context) error {
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

	req := &api.PostCliplistItemRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	cliplist, err := c.CliplistRepository.NewQuery(ctx.GetDB()).
		Where(&model.Cliplist{
			Status: model.CliplistStatusPublic,
		}).
		Find()
	if err != nil {
		if repository.NotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "cliplist was not found")
		}
		return err
	}
	if cliplist.UserID != app_context.GetUserID(ctx) {
		return echo.NewHTTPError(http.StatusForbidden, "cliplist is not yours")
	}

	cliplistContains, err := c.CliplistContainRepository.NewQuery(ctx.GetDB()).
		Where(&model.CliplistContain{
			CliplistID: cliplistID,
		}).FindAll()
	if err != nil {
		return err
	}
	if index > len(cliplistContains) {
		return echo.NewHTTPError(http.StatusBadRequest, "cliplist index was out of range")
	}

	_, err = c.ClipRepository.NewQuery(ctx.GetDB()).Where(&model.Clip{
		ID:     req.ClipID,
		Status: model.ClipStatusPublic,
	}).Find()
	if err != nil {
		if repository.NotFoundError(err) {
			return echo.NewHTTPError(http.StatusBadRequest, "clip was not found")
		}
		return err
	}

	cliplistContain := model.NewCliplistContain(
		cliplistID,
		index,
		req.ClipID,
	)
	err = c.CliplistContainRepository.InsertToList(ctx.GetDB(), cliplistContain)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, &api.PostCliplistItemResponse{
		CliplistID: cliplistID,
	})
}

func (c *CliplistController) DeleteCliplistItem(ctx context.Context) error {
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

	cliplist, err := c.CliplistRepository.NewQuery(ctx.GetDB()).
		Where(&model.Cliplist{
			Status: model.CliplistStatusPublic,
		}).
		Find()
	if err != nil {
		if repository.NotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "cliplist was not found")
		}
		return err
	}
	if cliplist.UserID != app_context.GetUserID(ctx) {
		return echo.NewHTTPError(http.StatusForbidden, "cliplist is not yours")
	}

	cliplistContains, err := c.CliplistContainRepository.NewQuery(ctx.GetDB()).
		Where(&model.CliplistContain{
			CliplistID: cliplistID,
		}).FindAll()
	if err != nil {
		return err
	}
	if index > len(cliplistContains) {
		return echo.NewHTTPError(http.StatusBadRequest, "cliplist index was out of range")
	}

	err = c.CliplistContainRepository.DeleteFromList(ctx.GetDB(), cliplistContains[index])
	if err != nil {
		return err
	}

	return nil
}
