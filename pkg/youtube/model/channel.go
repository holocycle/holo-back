package model

import "time"

type Channel struct {
	Kind    string
	ETag    string
	ID      string
	Snippet struct {
		Title       string
		Description string
		PublishedAt time.Time
		Thumbnails  struct {
			Default struct {
				URL    string
				Width  int
				Height int
			}
			Medium struct {
				URL    string
				Width  int
				Height int
			}
			High struct {
				URL    string
				Width  int
				Height int
			}
		}
	}
	ContentDetails struct {
		RelatedPlaylists struct {
			Likes            string
			Favorites        string
			Uploads          string
			WatchHistory     string
			WatchLater       string
			GooglePlusUserID string
		}
	}
	Statistics struct {
		ViewCount             JSONInt64
		CommentCount          JSONInt64
		SubscriberCount       JSONInt64
		HiddenSubscriberCount bool
		VideoCount            JSONInt64
	}
	BrandingSettings struct {
		Channel struct {
			Title                      string
			Description                string
			Keywords                   string
			DefaultTab                 string
			TrackingAnalyticsAccountID string
			ModerateComments           bool
			ShowRelatedChannels        bool
			ShowBrowseView             bool
			FeaturedChannelsTitle      string
			FeaturedChannelURLs        []string
			UnsubscribedTrailer        string
			ProfileColor               string
			Watch                      struct {
				TextColor          string
				BackgroundColor    string
				FeaturedPlaylistID string
			}
		}
		Image struct {
			BannerImageURL                    string
			BannerMobileImageURL              string
			BackgroundImageURL                LocalizedImageURL
			LargeBrandedBannerImageImapScript LocalizedImageURL
			LargeBrandedBannerImageURL        LocalizedImageURL
			SmallBrandedBannerImageImapScript LocalizedImageURL
			SmallBrandedBannerImageURL        LocalizedImageURL
			WatchIconImageURL                 string
			TrackingImageURL                  string
			BannerTabletLowImageURL           string
			BannerTabletImageURL              string
			BannerTabletHdImageURL            string
			BannerTabletExtraHdImageURL       string
			BannerMobileLowImageURL           string
			BannerMobileMediumHdImageURL      string
			BannerMobileHdImageURL            string
			BannerMobileExtraHdImageURL       string
			BannerTvImageURL                  string
			BannerExternalURL                 string
		}
		Hints []struct {
			Property string
			Value    string
		}
	}
}

type LocalizedImageURL struct {
	Default   string
	Localized []struct {
		Value    string
		Language string
	}
}

type ChannelListResponse struct {
	Kind          string
	ETag          string
	NextPageToken string
	PrevPageToken string
	PageInfo      struct {
		TotalResults   int
		ResultsPerPage int
	}
	Items []Channel
}
