package youtube

import (
	"encoding/json"
	"errors"

	"github.com/holocycle/holo-back/pkg/httpclient"
	"github.com/holocycle/holo-back/pkg/model"
	youtube_model "github.com/holocycle/holo-back/pkg/youtube/model"
)

func (c *ClientImpl) GetChannel(channelID string) (*model.Channel, error) {
	bytes, err := httpclient.GetRaw(c.APIURL.Channel, map[string]string{
		"key":  c.APIKey,
		"id":   channelID,
		"part": "snippet,statistics,brandingSettings",
	})
	if err != nil {
		return nil, newErr(err)
	}

	resp := &youtube_model.ChannelListResponse{}
	if err := json.Unmarshal(bytes, resp); err != nil {
		return nil, newErr(err)
	}
	if len(resp.Items) == 0 {
		return nil, newErr(errors.New("Channel is not found"))
	}

	ch := resp.Items[0]
	channel := model.NewChannel(
		ch.ID,
		ch.Snippet.Title,
		ch.Snippet.Description,
		ch.Snippet.Thumbnails.Default.URL,
		ch.Snippet.Thumbnails.Medium.URL,
		ch.Snippet.Thumbnails.High.URL,
		ch.BrandingSettings.Image.BannerTabletLowImageURL,
		ch.BrandingSettings.Image.BannerTabletImageURL,
		ch.BrandingSettings.Image.BannerTabletHdImageURL,
		ch.Statistics.ViewCount.Value,
		ch.Statistics.CommentCount.Value,
		ch.Statistics.SubscriberCount.Value,
		ch.Statistics.VideoCount.Value,
		&ch.Snippet.PublishedAt,
	)
	return channel, nil
}
