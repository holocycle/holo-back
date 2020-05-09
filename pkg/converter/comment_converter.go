package converter

import (
	"github.com/holocycle/holo-back/pkg/api"
	"github.com/holocycle/holo-back/pkg/model"
)

func ConvertToComment(comment *model.Comment) *api.Comment {
	return &api.Comment{
		ModelBase: api.ModelBase{
			Type: "Comment",
			ID:   comment.ID,
		},
		UserID:  comment.UserID,
		ClipID:  comment.ClipID,
		Content: comment.Content,
		User:    ConvertToUser(comment.User),
	}
}

func ConvertToComments(comments []*model.Comment) []*api.Comment {
	res := make([]*api.Comment, 0)
	for _, comment := range comments {
		res = append(res, ConvertToComment(comment))
	}
	return res
}
