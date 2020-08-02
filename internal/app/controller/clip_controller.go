package controller

import (
	"net/http"

	"github.com/holocycle/holo-back/internal/app/config"
	"github.com/holocycle/holo-back/pkg/api"
	app_context "github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/converter"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/holocycle/holo-back/pkg/repository"
	"github.com/holocycle/holo-back/pkg/youtube"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type ClipController struct {
	Config              *config.AppConfig
	RepositoryContainer *repository.Container
}

func NewClipController(config *config.AppConfig) *ClipController {
	return &ClipController{
		Config:              config,
		RepositoryContainer: repository.NewContainer(),
	}
}

func (c *ClipController) Register(e *echo.Echo) {
	get(e, "/clips", c.ListClips)
	post(e, "/clips", c.PostClip)

	get(e, "/clips/:clip_id", c.GetClip)
	put(e, "/clips/:clip_id", c.PutClip)
	delete(e, "/clips/:clip_id", c.DeleteClip)
}

func (c *ClipController) ListClips(ctx echo.Context) error {
	goCtx := ctx.Request().Context()
	log := app_context.GetLog(goCtx)

	req := &api.ListClipsRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	log.Debug("success to validate", zap.Any("req", req))

	query := c.RepositoryContainer.ClipRepository.NewQuery(goCtx).
		Where(&model.Clip{Status: model.ClipStatusPublic})

	// clip作成者を絞り込む
	createdBy := req.CreatedBy
	if createdBy != "" {
		query = query.Where(&model.Clip{UserID: createdBy})
	}

	// tag情報から絞り込みに利用する情報を取得する
	tags := req.Tags
	if len(tags) > 0 {
		query = query.JoinClipTaggedIn(tags)
	}

	query = query.JoinVideo().
		JoinFavorite()
	if req.Limit > 0 {
		query = query.Limit(req.Limit)
	}

	if req.OrderBy == "latest" {
		query = query.Latest()
	} else if req.OrderBy == "toprated" {
		query = query.TopRated()
	}

	// 絞り込み用のtagが指定されていた場合、すべてのtagが付与されているものを対象とする。
	if len(tags) > 0 {
		query.Having("COUNT(distinct clip_tagged.tag_id) = (?)", len(tags))
	}

	clips, err := query.FindAll()
	if err != nil {
		return err
	}
	log.Debug("success to retrieve clips")

	return ctx.JSON(http.StatusOK, &api.ListClipsResponse{
		Clips: converter.ConvertToClips(clips),
	})
}

func (c *ClipController) PostClip(ctx echo.Context) error {
	goCtx := ctx.Request().Context()
	log := app_context.GetLog(goCtx)

	req := &api.PostClipRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	log.Debug("success to validate", zap.Any("req", req))

	youtubeCli := youtube.New(&c.Config.YoutubeClient)
	video, err := youtubeCli.GetVideo(req.VideoID)
	if err != nil {
		return err
	}
	log.Debug("success to retireve video info from youtube", zap.Any("video", video))

	if req.BeginAt > video.Duration {
		log.Info("failed to validate duration",
			zap.Int("req.BeginAt", req.BeginAt),
			zap.Int("video.Duration", video.Duration),
		)
		return echo.NewHTTPError(http.StatusBadRequest, "begint_at is out of range")
	}
	if req.EndAt > video.Duration {
		log.Info("failed to validate duration",
			zap.Int("req.BeginAt", req.BeginAt),
			zap.Int("video.Duration", video.Duration),
		)
		return echo.NewHTTPError(http.StatusBadRequest, "end_at is out of range")
	}
	log.Debug("success to validate duration",
		zap.Int("req.BeginAt", req.BeginAt),
		zap.Int("req.EndAt", req.BeginAt),
		zap.Int("video.Duration", video.Duration),
	)

	if err := c.RepositoryContainer.VideoRepository.NewQuery(goCtx).Save(video); err != nil {
		log.Error("failed to save video", zap.Any("video", video))
		return err
	}
	log.Debug("success to save video", zap.Any("video", video))

	clip := model.NewClip(
		app_context.GetSession(goCtx).UserID,
		req.Title,
		req.Description,
		req.VideoID,
		req.BeginAt,
		req.EndAt,
		model.ClipStatusPublic,
	)
	if err := c.RepositoryContainer.ClipRepository.NewQuery(goCtx).Create(clip); err != nil {
		log.Error("failed to create clip", zap.Any("clip", clip))
		return err
	}
	log.Debug("success to create video", zap.Any("clip", clip))

	return ctx.JSON(http.StatusCreated, &api.PostClipResponse{
		ClipID: clip.ID,
	})
}

func (c *ClipController) GetClip(ctx echo.Context) error {
	goCtx := ctx.Request().Context()
	log := app_context.GetLog(goCtx)

	clipID := ctx.Param("clip_id")
	if clipID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify clip_id")
	}
	log.Debug("success to retrieve path parameter", zap.String("clipId", clipID))

	clip, err := c.RepositoryContainer.ClipRepository.NewQuery(goCtx).
		JoinVideo().
		JoinFavorite().
		Where(&model.Clip{ID: clipID, Status: model.ClipStatusPublic}).
		Find()
	if err != nil {
		if repository.NotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "clip was not found")
		}
		return err
	}
	log.Debug("success to retrieve clip", zap.Any("clip", clip))

	return ctx.JSON(http.StatusOK, &api.GetClipResponse{
		Clip: converter.ConvertToClip(clip, clip.Video, clip.Favorites),
	})
}

func (c *ClipController) PutClip(ctx echo.Context) error {
	goCtx := ctx.Request().Context()
	log := app_context.GetLog(goCtx)

	clipID := ctx.Param("clip_id")
	if clipID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify clip_id")
	}

	req := &api.PutClipRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	clip, err := c.RepositoryContainer.ClipRepository.NewQuery(goCtx).
		Where(&model.Clip{ID: clipID, Status: model.ClipStatusPublic}).
		Find()
	if err != nil {
		if repository.NotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "clip was not found")
		}
		return err
	}

	if app_context.GetSession(goCtx).UserID != clip.UserID {
		return echo.NewHTTPError(http.StatusForbidden, "clip is not yours")
	}
	log.Debug("success to validate parameters")

	clip.Title = req.Title
	clip.Description = req.Description
	clip.BeginAt = req.BeginAt
	clip.EndAt = req.EndAt
	if err := c.RepositoryContainer.ClipRepository.NewQuery(goCtx).Save(clip); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, &api.PutClipResponse{
		ClipID: clipID,
	})
}

func (c *ClipController) DeleteClip(ctx echo.Context) error {
	goCtx := ctx.Request().Context()
	log := app_context.GetLog(goCtx)

	clipID := ctx.Param("clip_id")
	if clipID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify clip_id")
	}

	clip, err := c.RepositoryContainer.ClipRepository.NewQuery(goCtx).
		Where(&model.Clip{ID: clipID, Status: model.ClipStatusPublic}).
		Find()
	if err != nil {
		if repository.NotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "clip was not found")
		}
		return err
	}

	if app_context.GetSession(goCtx).UserID != clip.UserID {
		return echo.NewHTTPError(http.StatusForbidden, "clip is not yours")
	}
	log.Debug("success to validate parameters")

	clip.Status = model.ClipStatusDeleted
	if err := c.RepositoryContainer.ClipRepository.NewQuery(goCtx).Save(clip); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, &api.DeleteClipRequest{})
}
