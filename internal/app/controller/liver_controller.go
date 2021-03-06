package controller

import (
	"net/http"
	"time"

	"github.com/holocycle/holo-back/internal/app/config"
	"github.com/holocycle/holo-back/pkg/api"
	"github.com/holocycle/holo-back/pkg/converter"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/holocycle/holo-back/pkg/repository"
	"github.com/holocycle/holo-back/pkg/youtube"
	"github.com/labstack/echo/v4"
)

type LiverController struct {
	Config              *config.AppConfig
	RepositoryContainer *repository.Container
}

func NewLiverController(config *config.AppConfig) *LiverController {
	return &LiverController{
		Config:              config,
		RepositoryContainer: repository.NewContainer(),
	}
}

func (c *LiverController) Register(e *echo.Echo) {
	get(e, "/livers", c.ListLivers)
	get(e, "/livers/:liver_id", c.GetLiver)
}

func (c *LiverController) ListLivers(ctx echo.Context) error {
	goCtx := ctx.Request().Context()
	livers, err := c.RepositoryContainer.LiverRepository.NewQuery(goCtx).JoinChannel().FindAll()
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
			if err := c.RepositoryContainer.ChannelRepository.NewQuery(goCtx).Save(channel); err != nil {
				return err
			}
		}

		livers, err = c.RepositoryContainer.LiverRepository.NewQuery(goCtx).JoinChannel().FindAll()
		if err != nil {
			return err
		}
	}

	return ctx.JSON(http.StatusOK, &api.ListLiversResponse{
		Livers: converter.ConvertToLivers(livers),
	})
}

func (c *LiverController) GetLiver(ctx echo.Context) error {
	goCtx := ctx.Request().Context()
	liverID := ctx.Param("liver_id")
	if liverID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify liver_id")
	}

	liver, err := c.RepositoryContainer.LiverRepository.NewQuery(goCtx).
		JoinChannel().Where(&model.Liver{ID: liverID}).Find()
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "liver was not found")
	}
	return ctx.JSON(http.StatusOK, &api.GetLiverResponse{
		Liver: converter.ConvertToLiver(liver, liver.Channel),
	})
}
