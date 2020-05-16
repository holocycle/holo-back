package controller

import (
	"net/http"
	"time"

	"github.com/holocycle/holo-back/internal/app/config"
	"github.com/holocycle/holo-back/pkg/api"
	"github.com/holocycle/holo-back/pkg/context"
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
	tx := ctx.GetDB()
	livers, err := c.LiverRepository.NewQuery(tx).JoinChannel().FindAll()
	if err != nil {
		return err
	}

	// TODO: extract as batch
	youtubeCli := youtube.New(&c.Config.YoutubeClient)

	curTime := time.Now()
	cacheDuration := 300 * time.Second
	for _, liver := range livers {
		if liver.Channel != nil &&
			curTime.Before(liver.Channel.UpdatedAt.Add(cacheDuration)) {
			continue
		}

		// there is no cache or cache is expired.
		channel, err := youtubeCli.GetChannel(liver.ChannelID)
		if err != nil {
			return err
		}
		liver.Channel = channel

		err = c.ChannelRepository.NewQuery(ctx.GetDB()).Save(channel)
		if err != nil {
			return err
		}
	}

	return ctx.JSON(http.StatusOK, &api.ListLiversResponse{
		Livers: converter.ConvertToLivers(livers),
	})
}

func (c *LiverController) GetLiver(ctx context.Context) error {
	liverID := ctx.Param("liver_id")
	if liverID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify liver_id")
	}

	tx := ctx.GetDB()
	liver, err := c.LiverRepository.NewQuery(tx).
		JoinChannel().Where(&model.Liver{ID: liverID}).Find()
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "liver was not found")
	}
	return ctx.JSON(http.StatusOK, &api.GetLiverResponse{
		Liver: converter.ConvertToLiver(liver, liver.Channel),
	})
}
