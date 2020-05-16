package api

import "time"

type Liver struct {
	ModelBase
	Name      string  `json:"name"`
	MainColor string  `json:"mainColor"`
	SubColor  string  `json:"subColor"`
	Channel   Channel `json:"channel"`
}

type Channel struct {
	ModelBase
	Title              string     `json:"title"`
	Description        string     `json:"description"`
	SmallThumbnailURL  string     `json:"smallThumbnailUrl"`
	MediumThumbnailURL string     `json:"mediumThumbnailUrl"`
	LargeThumbnailURL  string     `json:"largeThumbnailUrl"`
	SmallBannerURL     string     `json:"smallBannerUrl"`
	MediumBannerURL    string     `json:"mediumBannerUrl"`
	LargeBannerURL     string     `json:"largeBannerUrl"`
	ViewCount          int64      `json:"viewCount"`
	CommentCount       int64      `json:"commentCount"`
	SubscriberCount    int64      `json:"subscriberCount"`
	VideoCount         int64      `json:"videoCount"`
	PublishedAt        *time.Time `json:"publishedAt"`
}

type ListLiversRequest struct {
}

type ListLiversResponse struct {
	Livers []*Liver `json:"livers"`
}

type GetLiverRequest struct {
}

type GetLiverResponse struct {
	Liver *Liver `json:"liver"`
}

func LiverModels() []interface{} {
	return []interface{}{
		Liver{},
		Channel{},
		ListLiversRequest{},
		ListLiversResponse{},
		GetLiverRequest{},
		GetLiverResponse{},
	}
}
