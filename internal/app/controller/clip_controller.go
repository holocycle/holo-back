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
	"github.com/holocycle/holo-back/pkg/youtube"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type ClipController struct {
	Config          *config.AppConfig
	ClipRepository  repository.ClipRepository
	VideoRepository repository.VideoRepository
}

func NewClipController(config *config.AppConfig) *ClipController {
	return &ClipController{
		Config:          config,
		ClipRepository:  repository.NewClipRepository(),
		VideoRepository: repository.NewVideoRepository(),
	}
}

func (c *ClipController) Register(e *echo.Echo) {
	get(e, "/clips", c.ListClips)
	post(e, "/clips", c.PostClip)

	get(e, "/clips/:clip_id", c.GetClip)
	put(e, "/clips/:clip_id", c.PutClip)
	delete(e, "/clips/:clip_id", c.DeleteClip)
}

func (c *ClipController) ListClips(ctx context.Context) error {
	log := ctx.GetLog()

	req := &api.ListClipsRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	log.Debug("success to validate", zap.Any("req", req))

	query := c.ClipRepository.NewQuery(ctx.GetDB()).JoinVideo()
	if req.Limit > 0 {
		query = query.Limit(req.Limit)
	}
	if req.OrderBy == "latest" {
		query = query.Latest()
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

func (c *ClipController) PostClip(ctx context.Context) error {
	log := ctx.GetLog()

	req := &api.PostClipRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	log.Debug("success to validate", zap.Any("req", req))

	youtubeCli := youtube.New(&c.Config.YoutubeClient)
	video, err := youtubeCli.GetVideo(req.VideoID)
	if err != nil {
		return err // FIXME
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

	tx := ctx.GetDB()
	if err := c.VideoRepository.NewQuery(tx).Save(video); err != nil {
		log.Error("failed to save video", zap.Any("video", video))
		return err
	}
	log.Debug("success to save video", zap.Any("video", video))

	clip := model.NewClip(
		app_context.GetSession(ctx).UserID,
		req.Title,
		req.Description,
		req.VideoID,
		req.BeginAt,
		req.EndAt,
	)
	if err := c.ClipRepository.NewQuery(tx).Create(clip); err != nil {
		log.Error("failed to create clip", zap.Any("clip", clip))
		return err
	}
	log.Debug("success to create video", zap.Any("clip", clip))

	return ctx.JSON(http.StatusCreated, &api.PostClipResponse{
		ClipID: clip.ID,
	})
}

func (c *ClipController) GetClip(ctx context.Context) error {
	log := ctx.GetLog()

	clipID := ctx.Param("clip_id")
	if clipID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify clip_id")
	}
	log.Debug("success to retrieve path parameter", zap.String("clipId", clipID))

	clip, err := c.ClipRepository.NewQuery(ctx.GetDB()).JoinVideo().
		Where(&model.Clip{ID: clipID}).Find()
	if err != nil {
		if repository.NotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "clip was not found")
		}
		return err
	}
	log.Debug("success to retrieve clip", zap.Any("clip", clip))

	return ctx.JSON(http.StatusOK, &api.GetClipResponse{
		Clip: converter.ConvertToClip(clip, clip.Video),
	})
}

func (c *ClipController) PutClip(ctx context.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

func (c *ClipController) DeleteClip(ctx context.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}
