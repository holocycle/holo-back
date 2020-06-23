package controller

import (
	"net/http"
	"time"

	"github.com/holocycle/holo-back/internal/app/config"
	"github.com/holocycle/holo-back/pkg/api"
	"github.com/holocycle/holo-back/pkg/context"
	app_context2 "github.com/holocycle/holo-back/pkg/context2"
	"github.com/holocycle/holo-back/pkg/converter"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/holocycle/holo-back/pkg/repository"
	"github.com/holocycle/holo-back/pkg/youtube"
	"github.com/labstack/echo/v4"
)

type LiverController struct {
	Config            *config.AppConfig
	LiverRepository   repository.LiverRepository
	ChannelRepository repository.ChannelRepository
}

func NewLiverController(config *config.AppConfig) *LiverController {
	return &LiverController{
		Config:            config,
		LiverRepository:   repository.NewLiverRepository(),
		ChannelRepository: repository.NewChannelRepository(),
	}
}

func (c *LiverController) Register(e *echo.Echo) {
	get(e, "/livers", c.ListLivers)
	get(e, "/livers/:liver_id", c.GetLiver)
}

func (c *LiverController) ListLivers(ctx context.Context) error {
	goCtx := app_context2.FromEchoContext(ctx)
	livers, err := c.LiverRepository.NewQuery(goCtx).JoinChannel().FindAll()
	if err != nil {
		return err
	}

	curTime := time.Now()
	cacheDuration := 300 * time.Second
	expiredChannelIDs := make([]string, 0)
	for _, liver := range livers {
		if liver.Channel == nil || curTime.After(liver.Channel.UpdatedAt.Add(cacheDuration)) {
			expiredChannelIDs = append(expiredChannelIDs, liver.ChannelID)
		}
	}

	if len(expiredChannelIDs) > 0 {
		youtubeCli := youtube.New(&c.Config.YoutubeClient)
		channels, err := youtubeCli.ListChannels(expiredChannelIDs)
		if err != nil {
			return err
		}

		for _, channel := range channels {
			if err := c.ChannelRepository.NewQuery(goCtx).Save(channel); err != nil {
				return err
			}
		}

		livers, err = c.LiverRepository.NewQuery(goCtx).JoinChannel().FindAll()
		if err != nil {
			return err
		}
	}

	return ctx.JSON(http.StatusOK, &api.ListLiversResponse{
		Livers: converter.ConvertToLivers(livers),
	})
}

func (c *LiverController) GetLiver(ctx context.Context) error {
	goCtx := app_context2.FromEchoContext(ctx)
	liverID := ctx.Param("liver_id")
	if liverID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify liver_id")
	}

	liver, err := c.LiverRepository.NewQuery(goCtx).
		JoinChannel().Where(&model.Liver{ID: liverID}).Find()
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "liver was not found")
	}
	return ctx.JSON(http.StatusOK, &api.GetLiverResponse{
		Liver: converter.ConvertToLiver(liver, liver.Channel),
	})
}
