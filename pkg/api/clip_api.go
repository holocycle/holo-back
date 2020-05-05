package api

type Clip struct {
	ModelBase
	Title       string `json:"title"`
	Description string `json:"description"`
	VideoID     string `json:"videoId"`
	BeginAt     int    `json:"beginAt"`
	EndAt       int    `json:"endAt"`
	Video       *Video `json:"video"`
}

type ListClipsRequest struct {
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

func ClipModels() []interface{} {
	return []interface{}{
		Clip{},
		PostClipRequest{},
		PostClipResponse{},
		GetClipRequest{},
		GetClipResponse{},
	}
}
