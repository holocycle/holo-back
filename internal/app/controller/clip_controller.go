package controller

import (
	"errors"
	"net/http"

	app_context "github.com/holocycle/holo-back/internal/app/context"
	"github.com/holocycle/holo-back/pkg/api"
	"github.com/holocycle/holo-back/pkg/context"
	"github.com/holocycle/holo-back/pkg/converter"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/holocycle/holo-back/pkg/repository"
	"github.com/holocycle/holo-back/pkg/youtube_client"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func RegisterClipController(e *echo.Echo) {
	e.GET("/clips", ListClips)
	e.POST("/clips", PostClip)

	e.GET("/clips/:clip_id", GetClip)
	e.PUT("/clips/:clip_id", PutClip)
	e.DELETE("/clips/:clip_id", DeleteClip)
}

func ListClips(c echo.Context) error {
	ctx := c.(context.Context)
	log := ctx.GetLog()

	req := &api.ListClipsRequest{}
	if err := ctx.Bind(req); err != nil {
		return err
	}
	if err := ctx.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	log.Info("success to validate", zap.Any("req", req))

	cond := &repository.ClipCondition{}
	cond.Limit = req.Limit
	if req.OrderBy == "latest" {
		cond.OrderBy = repository.RecentlyCreated
	}

	clipRepo := repository.NewClipRepository(ctx)
	clips, err := clipRepo.FindAll(cond)
	if err != nil {
		return err
	}
	log.Info("success to retrieve clips")

	videoRepo := repository.NewVideoRepository(ctx)
	videos, err := videoRepo.FindAll(&model.Video{})
	if err != nil {
		return err
	}
	log.Info("success to retrieve videos")

	videoMap := make(map[string]*model.Video)
	for _, video := range videos {
		videoMap[video.ID] = video
	}

	res := make([]*api.Clip, 0)
	for _, clip := range clips {
		video, ok := videoMap[clip.VideoID]
		if !ok {
			log.Error("found clip but video was not found", zap.String("videoID", clip.VideoID))
			return errors.New("found clip but video was not found")
		}
		res = append(res, converter.ConvertToClip(clip, video))
	}

	return ctx.JSON(http.StatusOK, &api.ListClipsResponse{
		Clips: res,
	})
}

func PostClip(c echo.Context) error {
	ctx := c.(context.Context)
	log := ctx.GetLog()
	cfg := app_context.GetConfig(ctx)

	req := &api.PostClipRequest{}
	if err := ctx.Bind(req); err != nil {
		return err
	}

	if err := ctx.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	log.Info("success to validate", zap.Any("req", req))

	youtubeCli := youtube_client.New(&cfg.YoutubeClient)
	video, err := youtubeCli.GetVideo(req.VideoID)
	if err != nil {
		return err // FIXME
	}
	log.Info("success to retireve video info from youtube", zap.Any("video", video))

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
	log.Info("success to validate duration",
		zap.Int("req.BeginAt", req.BeginAt),
		zap.Int("req.EndAt", req.BeginAt),
		zap.Int("video.Duration", video.Duration),
	)

	videoRepo := repository.NewVideoRepository(ctx)
	if err := videoRepo.Save(video); err != nil {
		log.Error("failed to save video", zap.Any("video", video))
		return err
	}
	log.Info("success to save video", zap.Any("video", video))

	clip := model.NewClip(
		app_context.GetSession(ctx).UserID,
		req.Title,
		req.Description,
		req.VideoID,
		req.BeginAt,
		req.EndAt,
	)
	clipRepo := repository.NewClipRepository(ctx)
	if err := clipRepo.Create(clip); err != nil {
		log.Error("failed to create clip", zap.Any("clip", clip))
		return err
	}
	log.Info("success to create video", zap.Any("clip", clip))

	return ctx.JSON(http.StatusCreated, &api.PostClipResponse{
		ClipID: clip.ID,
	})
}

func GetClip(c echo.Context) error {
	ctx := c.(context.Context)
	log := ctx.GetLog()

	clipId := ctx.Param("clip_id")
	if clipId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify clip_id")
	}
	log.Info("success to retrieve path parameter", zap.String("clipId", clipId))

	clipRepo := repository.NewClipRepository(ctx)
	clip, err := clipRepo.FindBy(&model.Clip{ID: clipId})
	if err != nil {
		if repository.NotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "clip was not found")
		}
		return err
	}
	log.Info("success to retrieve clip", zap.Any("clip", clip))

	videoRepo := repository.NewVideoRepository(ctx)
	video, err := videoRepo.FindBy(&model.Video{ID: clip.VideoID})
	if err != nil {
		if repository.NotFoundError(err) {
			log.Error("no video for clip", zap.Any("clip", clip))
			return err
		}
		return err
	}
	log.Info("success to retrieve video", zap.Any("video", video))

	return ctx.JSON(http.StatusOK, &api.GetClipResponse{
		Clip: converter.ConvertToClip(clip, video),
	})
}

func PutClip(c echo.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

func DeleteClip(c echo.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}
