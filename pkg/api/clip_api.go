package api

type Clip struct {
	ModelBase
	Title         string `json:"title"`
	Description   string `json:"description"`
	BeginAt       int    `json:"beginAt"`
	EndAt         int    `json:"endAt"`
	FavoriteCount int    `json:"favoriteCount"`
	Video         *Video `json:"video"`
}

type ListClipsRequest struct {
	Limit     int      `json:"limit"   validate:"min=0,max=100"`
	OrderBy   string   `json:"orderBy" validate:"oneof=any latest toprated"`
	Tags      []string `json:"tags"`
	CreatedBy string   `json:"createdBy"`
}

type ListClipsResponse struct {
	Clips []*Clip `json:"clips"`
}

type PostClipRequest struct {
	VideoID     string `json:"videoId"     validate:"required,max=64"`
	Title       string `json:"title"       validate:"required,max=255"`
	Description string `json:"description" validate:"required"`
	BeginAt     int    `json:"beginAt"     validate:"gte=0"`
	EndAt       int    `json:"endAt"       validate:"gtfield=BeginAt"`
}

type PostClipResponse struct {
	ClipID string `json:"clipId"`
}

type GetClipRequest struct {
}

type GetClipResponse struct {
	Clip *Clip `json:"clip"`
}

type PutClipRequest struct {
	Title       string `json:"title"       validate:"required,max=255"`
	Description string `json:"description" validate:"required"`
	BeginAt     int    `json:"beginAt"     validate:"gte=0"`
	EndAt       int    `json:"endAt"       validate:"gtfield=BeginAt"`
}

type PutClipResponse struct {
	ClipID string `json:"clipId"`
}

type DeleteClipRequest struct {
}

type DeleteClipResponse struct {
}

func ClipModels() []interface{} {
	return []interface{}{
		Clip{},
		ListClipsRequest{},
		ListClipsResponse{},
		PostClipRequest{},
		PostClipResponse{},
		GetClipRequest{},
		GetClipResponse{},
		PutClipRequest{},
		PutClipResponse{},
		DeleteClipRequest{},
		DeleteClipResponse{},
	}
}
