package youtube_client

import (
	"encoding/json"
	"errors"

	"github.com/holocycle/holo-back/pkg/http_client"
	"github.com/holocycle/holo-back/pkg/model"
	youtube_model "github.com/holocycle/holo-back/pkg/youtube_client/model"
)

func (c *Client) GetVideo(videoID string) (*model.Video, error) {
	bytes, err := http_client.GetRaw(c.APIURL.Video, map[string]string{
		"key":  c.APIKey,
		"id":   videoID,
		"part": "snippet,contentDetails",
	})
	if err != nil {
		return nil, err
	}

	resp := &youtube_model.VideoListResponse{}
	if err := json.Unmarshal(bytes, resp); err != nil {
		return nil, err
	}

	if len(resp.Items) == 0 {
		return nil, errors.New("Video is not found")
	}

	v := resp.Items[0]
	duration := v.ContentDetails.Duration.TH*3600 +
		v.ContentDetails.Duration.TM*60 + v.ContentDetails.Duration.TS // FIXME
	video := model.NewVideo(
		v.ID,
		v.Snippet.ChannelID,
		v.Snippet.Title,
		v.Snippet.Description,
		duration,
		v.Snippet.Thumbnails.Default.URL,
		v.Snippet.Thumbnails.Medium.URL,
		v.Snippet.Thumbnails.Large.URL,
		&v.Snippet.PublishedAt,
	)

	return video, nil
}
