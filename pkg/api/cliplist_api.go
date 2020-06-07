package api

type Cliplist struct {
	ModelBase
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Length      int           `json:"length"`
	FirstItem   *CliplistItem `json:"firstItem"`
}

type CliplistItem struct {
	Clip
	Available bool `json:"available"`
}

type ListCliplistsRequest struct {
	Limit   int    `json:"limit"   validate:"min=0,max=100"`
	OrderBy string `json:"orderBy" validate:"oneof=any"`
}

type ListCliplistsResponse struct {
	Cliplists []*Cliplist `json:"cliplists"`
}

type GetCliplistRequest struct {
	Page        int `json:"page"        validate:"min=0"`
	ItemPerPage int `json:"itemPerPage" validate:"min=1,max=100"`
}

type GetCliplistResponse struct {
	Cliplist      *Cliplist       `json:"cliplist"`
	PageInfo      *PageInfo       `json:"pageInfo"`
	CliplistItems []*CliplistItem `json:"cliplistItems"`
}

type PostCliplistRequest struct {
	Title       string `json:"title"       validate:"min=0,max=255"`
	Description string `json:"description" validate:"min=0"`
}

type PostCliplistResponse struct {
	CliplistID string `json:"cliplistId"`
}

type PutCliplistRequest struct {
	Title       string `json:"title"       validate:"min=0,max=255"`
	Description string `json:"description" validate:"min=0"`
}

type PutCliplistResponse struct {
	CliplistID string `json:"cliplistId"`
}

type DeleteCliplistRequest struct {
}

type DeleteCliplistResponse struct {
}

type GetCliplistItemRequest struct {
}

type GetCliplistItemResponse struct {
	CliplistItem *CliplistItem `json:"cliplistItem"`
}

type PostCliplistItemRequest struct {
	ClipID string `json:"clipId" validate:"min=:,max=64"`
}

type PostCliplistItemResponse struct {
	CliplistID string `json:"cliplistId"`
}

type DeleteCliplistItemRequest struct {
}

type DeleteCliplistItemResponse struct {
	CliplistID string `json:"cliplistId"`
}

func CliplistModels() []interface{} {
	return []interface{}{
		Cliplist{},
		CliplistItem{},
		ListCliplistsRequest{},
		ListCliplistsResponse{},
		GetCliplistRequest{},
		GetCliplistResponse{},
		PostCliplistRequest{},
		PostCliplistResponse{},
		PutCliplistRequest{},
		PutCliplistResponse{},
		DeleteCliplistRequest{},
		DeleteCliplistResponse{},
		GetCliplistItemRequest{},
		GetCliplistItemResponse{},
		PostCliplistItemRequest{},
		PostCliplistItemResponse{},
		DeleteCliplistItemRequest{},
		DeleteCliplistItemResponse{},
	}
}
