package api

type Comment struct {
	ModelBase
	UserID  string `json:"userId"`
	ClipID  string `json:"clipId"`
	Content string `json:"content"`
	User    *User  `json:"user"`
}

type ListCommentsRequest struct {
	Limit   int    `json:"limit"   validate:"min=0,max=100"`
	OrderBy string `json:"orderBy" validate:"oneof=any latest"`
}

type ListCommentsResponse struct {
	Comments []*Comment `json:"comments"`
}

type GetCommentRequest struct {
}

type GetCommentResponse struct {
	Comment *Comment `json:"comment"`
}

type PostCommentRequest struct {
	Content string `json:"content"`
}

type PostCommentResponse struct {
	CommentID string `json:"commentId"`
}

type DeleteCommentRequest struct {
}

type DeleteCommentResponse struct {
}

func CommentModels() []interface{} {
	return []interface{}{
		Comment{},
		ListCommentsRequest{},
		ListCommentsResponse{},
		GetCommentRequest{},
		GetCommentResponse{},
		PostCommentRequest{},
		PostCommentResponse{},
		DeleteCommentRequest{},
		DeleteCommentResponse{},
	}
}
