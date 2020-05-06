package controller

import (
	"net/http"

	"github.com/holocycle/holo-back/internal/app/config"
	"github.com/holocycle/holo-back/pkg/api"
	"github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/converter"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/holocycle/holo-back/pkg/repository"
	"github.com/labstack/echo/v4"
)

type TagController struct {
	Config        *config.AppConfig
	TagRepository repository.TagRepository
}

func NewTagController(config *config.AppConfig) *TagController {
	return &TagController{
		Config:        config,
		TagRepository: repository.NewTagRepository(),
	}
}

func (c *TagController) Register(e *echo.Echo) {
	get(e, "/tags", c.ListTags)
	get(e, "/tags/:tag_id", c.GetTag)
	put(e, "/tags", c.PutTag)
}

func (c *TagController) ListTags(ctx context.Context) error {
	req := &api.ListTagsRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	tx := ctx.GetDB()
	tags, err := c.TagRepository.FindAll(tx, &repository.TagCondition{})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, &api.ListTagsResponse{
		Tags: converter.ConvertToTags(tags),
	})
}

func (c *TagController) GetTag(ctx context.Context) error {
	req := &api.GetTagRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	tagID := ctx.Param("tag_id")
	if tagID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify tag_id")
	}

	tx := ctx.GetDB()
	tag, err := c.TagRepository.FindBy(tx, &repository.TagCondition{ID: tagID})
	if err != nil {
		if repository.NotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "tag is not found")
		}
		return err
	}
	return ctx.JSON(http.StatusOK, &api.GetTagResponse{
		Tag: converter.ConvertToTag(tag),
	})
}

func (c *TagController) PutTag(ctx context.Context) error {
	req := &api.PutTagRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	tx := ctx.GetDB()
	tag, err := c.TagRepository.FindBy(tx, &repository.TagCondition{Name: req.Name})
	if err != nil && !repository.NotFoundError(err) {
		return err
	}
	if tag != nil {
		return ctx.JSON(http.StatusConflict, &api.PutTagResponse{
			TagID: tag.ID,
		})
	}

	tag = model.NewTag(req.Name, req.Color)
	if err := c.TagRepository.Save(tx, tag); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, &api.PutTagResponse{
		TagID: tag.ID,
	})
}
