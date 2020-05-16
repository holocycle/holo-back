package youtube

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/holocycle/holo-back/pkg/httpclient"
	"github.com/holocycle/holo-back/pkg/model"
	youtube_model "github.com/holocycle/holo-back/pkg/youtube/model"
)

func (c *ClientImpl) GetChannel(channelID string) (*model.Channel, error) {
	channels, err := c.ListChannels([]string{channelID})
	if err != nil {
		return nil, err
	}
	return channels[0], nil
}

func (c *ClientImpl) ListChannels(channelIDs []string) ([]*model.Channel, error) {
	// TODO: paging
	bytes, err := httpclient.GetRaw(c.APIURL.Channel, map[string]string{
		"key":  c.APIKey,
		"id":   strings.Join(channelIDs, ","),
		"part": "snippet,statistics,brandingSettings",
	})
	if err != nil {
		return nil, newErr(err)
	}

	resp := &youtube_model.ChannelListResponse{}
	if err := json.Unmarshal(bytes, resp); err != nil {
		return nil, newErr(err)
	}
	if len(resp.Items) < len(channelIDs) {
		return nil, newErr(errors.New("Channel is not found"))
	}

	channels := make([]*model.Channel, 0)
	for _, item := range resp.Items {
		channel := model.NewChannel(
			item.ID,
			item.Snippet.Title,
			item.Snippet.Description,
			item.Snippet.Thumbnails.Default.URL,
			item.Snippet.Thumbnails.Medium.URL,
			item.Snippet.Thumbnails.High.URL,
			item.BrandingSettings.Image.BannerTabletLowImageURL,
			item.BrandingSettings.Image.BannerTabletImageURL,
			item.BrandingSettings.Image.BannerTabletHdImageURL,
			item.Statistics.ViewCount.Value,
			item.Statistics.CommentCount.Value,
			item.Statistics.SubscriberCount.Value,
			item.Statistics.VideoCount.Value,
			&item.Snippet.PublishedAt,
		)
		channels = append(channels, channel)
	}

	return channels, nil
}
