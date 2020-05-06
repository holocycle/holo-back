package controller

import (
	"net/http"

	"github.com/holocycle/holo-back/internal/app/config"
	app_context "github.com/holocycle/holo-back/internal/app/context"
	"github.com/holocycle/holo-back/pkg/api"
	"github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/converter"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/holocycle/holo-back/pkg/repository"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type TagController struct {
	Config            *config.AppConfig
	TagRepository     repository.TagRepository
	ClipTagRepository repository.ClipTagRepository
}

func NewTagController(config *config.AppConfig) *TagController {
	return &TagController{
		Config:            config,
		TagRepository:     repository.NewTagRepository(),
		ClipTagRepository: repository.NewClipTagRepository(),
	}
}

func (c *TagController) Register(e *echo.Echo) {
	get(e, "/tags", c.ListTags)
	get(e, "/tags/:tag_id", c.GetTag)
	put(e, "/tags", c.PutTag)
	get(e, "/clips/:clip_id/tags", c.ListTagsOnClip)
	put(e, "/clips/:clip_id/tags/:tag_id", c.PutTagOnClip)
	delete(e, "/clips/:clip_id/tags/:tag_id", c.DeleteTagOnClip)
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

	return ctx.JSON(http.StatusCreated, &api.PutTagResponse{
		TagID: tag.ID,
	})
}

func (c *TagController) ListTagsOnClip(ctx context.Context) error {
	log := ctx.GetLog()

	req := &api.ListTagsOnClipRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	clipID := ctx.Param("clip_id")
	if clipID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify clip_id")
	}
	log.Debug("success to validate", zap.String("clipID", clipID))

	clipRepo := repository.NewClipRepository(ctx)
	clip, err := clipRepo.FindBy(&model.Clip{ID: clipID})
	if err != nil {
		if repository.NotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "clip is not found")
		}
		return err
	}
	log.Debug("success to retieve Clip", zap.Any("clip", clip))

	tx := ctx.GetDB()
	clipTags, err := c.ClipTagRepository.NewQuery(tx).
		JoinTag().Where(&model.ClipTag{ClipID: clipID}).FindAll()
	if err != nil {
		return err
	}
	log.Debug("success to retieve ClipTags", zap.Any("clipTags", clipTags))

	return ctx.JSON(http.StatusOK, &api.ListTagsOnClipResponse{
		ClipID: clip.ID,
		Tags:   converter.ConvertClipTagsToTags(clipTags),
	})
}

func (c *TagController) PutTagOnClip(ctx context.Context) error {
	log := ctx.GetLog()

	req := &api.PutTagOnClipRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	clipID := ctx.Param("clip_id")
	if clipID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify clip_id")
	}
	tagID := ctx.Param("tag_id")
	if tagID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify tag_id")
	}
	log.Debug("success to validate request", zap.String("clipID", clipID), zap.String("tagID", tagID))

	_, err := repository.NewClipRepository(ctx).FindBy(&model.Clip{ID: clipID})
	if err != nil {
		if repository.NotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "clip was not found")
		}
		return err
	}
	log.Debug("success to valipdate clipID", zap.String("clipID", clipID))

	tx := ctx.GetDB()
	_, err = c.TagRepository.FindBy(tx, &repository.TagCondition{ID: tagID})
	if err != nil {
		if repository.NotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "tag was not found")
		}
		return err
	}
	log.Debug("success to valipdate tagID", zap.String("tagID", tagID))

	clipTag, err := c.ClipTagRepository.NewQuery(tx).Where(&model.ClipTag{
		ClipID: clipID,
		TagID:  tagID,
	}).Find()
	if err != nil && !repository.NotFoundError(err) {
		return err
	}
	if clipTag != nil {
		return ctx.JSON(http.StatusConflict, &api.PutTagOnClipResponse{})
	}

	clipTag = model.NewClipTag(
		app_context.GetSession(ctx).UserID,
		clipID,
		tagID,
	)
	if err := c.ClipTagRepository.NewQuery(tx).Create(clipTag); err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, &api.PutTagOnClipResponse{})
}

func (c *TagController) DeleteTagOnClip(ctx context.Context) error {
	log := ctx.GetLog()

	req := &api.DeleteTagOnClipRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	clipID := ctx.Param("clip_id")
	if clipID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify clip_id")
	}
	tagID := ctx.Param("tag_id")
	if tagID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify tag_id")
	}
	log.Debug("success to validate request", zap.String("clipID", clipID), zap.String("tagID", tagID))

	tx := ctx.GetDB()
	rows, err := c.ClipTagRepository.NewQuery(tx).Where(&model.ClipTag{
		ClipID: clipID,
		TagID:  tagID,
	}).Delete()
	if err != nil {
		return err
	}
	if rows == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "tag on clip was not found")
	}

	return ctx.JSON(http.StatusOK, &api.DeleteTagOnClipResponse{})
}
