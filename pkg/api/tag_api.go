package api

type Tag struct {
	ModelBase
	Name  string `json:"name"`
	Color string `json:"color"`
}

type ListTagsRequest struct {
}

type ListTagsResponse struct {
	Tags []*Tag `json:"tags"`
}

type GetTagRequest struct {
}

type GetTagResponse struct {
	Tag *Tag `json:"tag"`
}

type PutTagRequest struct {
	Name  string `json:"name"  validate:"required"`
	Color string `json:"color" validate:"required,hexcolor"`
}

type PutTagResponse struct {
	TagID string `json:"tagId"`
}

func TagModels() []interface{} {
	return []interface{}{
		Tag{},
		ListTagsRequest{},
		ListTagsResponse{},
		GetTagRequest{},
		GetTagResponse{},
		PutTagRequest{},
		PutTagResponse{},
	}
}
