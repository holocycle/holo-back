package controller

import (
	"net/http"

	"github.com/holocycle/holo-back/internal/app/config"
	"github.com/holocycle/holo-back/pkg/context"
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
	return nil
}

func PostClip(c echo.Context) error {
	ctx := c.(context.Context)
	log := ctx.GetLog()
	cfg := ctx.Get("config").(*config.AppConfig)

	type Form struct {
		VideoID     string `json:"videoId"     validate:"required,max=64"`
		Title       string `json:"title"       validate:"required,max=255"`
		Description string `json:"description" validate:"required"`
		BeginAt     int    `json:"beginAt"     validate:"gte=0"`
		EndAt       int    `json:"endAt"       validate:"gtfield=BeginAt"`
	}

	form := &Form{}
	if err := ctx.Bind(form); err != nil {
		return err
	}

	if err := ctx.Validate(form); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	log.Info("success to validate", zap.Any("form", form))

	youtubeCli := youtube_client.New(&cfg.YoutubeClient)
	video, err := youtubeCli.GetVideo(form.VideoID)
	if err != nil {
		return err // FIXME
	}
	log.Info("success to retireve video info from youtube", zap.Any("video", video))

	if form.BeginAt > video.Duration {
		log.Info("failed to validate duration",
			zap.Int("form.BeginAt", form.BeginAt),
			zap.Int("video.Duration", video.Duration),
		)
		return echo.NewHTTPError(http.StatusBadRequest, "begint_at is out of range")
	}
	if form.EndAt > video.Duration {
		log.Info("failed to validate duration",
			zap.Int("form.BeginAt", form.BeginAt),
			zap.Int("video.Duration", video.Duration),
		)
		return echo.NewHTTPError(http.StatusBadRequest, "end_at is out of range")
	}
	log.Info("success to validate duration",
		zap.Int("form.BeginAt", form.BeginAt),
		zap.Int("form.EndAt", form.BeginAt),
		zap.Int("video.Duration", video.Duration),
	)

	videoRepo := repository.NewVideoRepository(ctx)
	if err := videoRepo.Save(video); err != nil {
		log.Error("failed to save video", zap.Any("video", video))
		return err
	}
	log.Info("success to save video", zap.Any("video", video))

	userID := "hoge" // FIXME
	clip := model.NewClip(
		userID,
		form.Title,
		form.Description,
		form.VideoID,
		form.BeginAt,
		form.EndAt,
	)
	clipRepo := repository.NewClipRepository(ctx)
	if err := clipRepo.Create(clip); err != nil {
		log.Error("failed to create clip", zap.Any("clip", clip))
		return err
	}
	log.Info("success to create video", zap.Any("clip", clip))

	return ctx.JSON(http.StatusCreated, clip)
}

func GetClip(c echo.Context) error {
	return nil
}

func PutClip(c echo.Context) error {
	return nil
}

func DeleteClip(c echo.Context) error {
	return nil
}
