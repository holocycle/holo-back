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
	Config               *config.AppConfig
	TagRepository        repository.TagRepository
	ClipRepository       repository.ClipRepository
	ClipTaggedRepository repository.ClipTaggedRepository
}

func NewTagController(config *config.AppConfig) *TagController {
	return &TagController{
		Config:               config,
		TagRepository:        repository.NewTagRepository(),
		ClipRepository:       repository.NewClipRepository(),
		ClipTaggedRepository: repository.NewClipTaggedRepository(),
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
	tags, err := c.TagRepository.NewQuery(tx).FindAll()
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
	tag, err := c.TagRepository.NewQuery(tx).Where(&model.Tag{ID: tagID}).Find()
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
	tag, err := c.TagRepository.NewQuery(tx).Where(&model.Tag{Name: req.Name}).Find()
	if err != nil && !repository.NotFoundError(err) {
		return err
	}
	if tag != nil {
		return ctx.JSON(http.StatusConflict, &api.PutTagResponse{
			TagID: tag.ID,
		})
	}

	tag = model.NewTag(req.Name, req.Color)
	if err := c.TagRepository.NewQuery(tx).Save(tag); err != nil {
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

	tx := ctx.GetDB()
	clip, err := c.ClipRepository.NewQuery(tx).
		Where(&model.Clip{ID: clipID, Status: model.CLIP_PUBLIC}).
		Find()
	if err != nil {
		if repository.NotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "clip is not found")
		}
		return err
	}
	log.Debug("success to retieve Clip", zap.Any("clip", clip))

	clipTagged, err := c.ClipTaggedRepository.NewQuery(tx).
		JoinTag().Where(&model.ClipTagged{ClipID: clipID}).FindAll()
	if err != nil {
		return err
	}
	log.Debug("success to retieve ClipTagged", zap.Any("clipTagged", clipTagged))

	return ctx.JSON(http.StatusOK, &api.ListTagsOnClipResponse{
		ClipID: clip.ID,
		Tags:   converter.ConvertClipTaggedToTags(clipTagged),
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
	log.Debug("success to validate request",
		zap.String("clipID", clipID),
		zap.String("tagID", tagID))

	tx := ctx.GetDB()
	_, err := repository.NewClipRepository().NewQuery(tx).
		Where(&model.Clip{ID: clipID, Status: model.CLIP_PUBLIC}).
		Find()
	if err != nil {
		if repository.NotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "clip was not found")
		}
		return err
	}
	log.Debug("success to valipdate clipID", zap.String("clipID", clipID))

	_, err = c.TagRepository.NewQuery(tx).Where(&model.Tag{ID: tagID}).Find()
	if err != nil {
		if repository.NotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "tag was not found")
		}
		return err
	}
	log.Debug("success to valipdate tagID", zap.String("tagID", tagID))

	cond := &model.ClipTagged{
		ClipID: clipID,
		TagID:  tagID,
	}
	clipTagged, err := c.ClipTaggedRepository.NewQuery(tx).Where(cond).Find()
	if err != nil && !repository.NotFoundError(err) {
		return err
	}
	if clipTagged != nil {
		return ctx.JSON(http.StatusConflict, &api.PutTagOnClipResponse{})
	}

	clipTagged = model.NewClipTagged(
		clipID,
		tagID,
		app_context.GetSession(ctx).UserID,
	)
	if err := c.ClipTaggedRepository.NewQuery(tx).Create(clipTagged); err != nil {
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

	tx := ctx.GetDB()
	_, err := repository.NewClipRepository().NewQuery(tx).
		Where(&model.Clip{ID: clipID, Status: model.CLIP_PUBLIC}).
		Find()
	if err != nil {
		if repository.NotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "clip was not found")
		}
		return err
	}
	log.Debug("success to validate request", zap.String("clipID", clipID), zap.String("tagID", tagID))

	cond := &model.ClipTagged{
		ClipID: clipID,
		TagID:  tagID,
	}
	rows, err := c.ClipTaggedRepository.NewQuery(tx).Where(cond).Delete()
	if err != nil {
		return err
	}
	if rows == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "tag on clip was not found")
	}

	return ctx.JSON(http.StatusOK, &api.DeleteTagOnClipResponse{})
}
