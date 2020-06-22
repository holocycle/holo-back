package model

import "time"

type Comment struct {
	ID        string
	UserID    string
	ClipID    string
	Content   string
	CreatedAt *time.Time
	UpdatedAt *time.Time

	User *User
}

func NewComment(
	userID,
	clipID,
	content string,
) *Comment {
	return &Comment{
		ID:      GetIDGenerator().New(),
		UserID:  userID,
		ClipID:  clipID,
		Content: content,
	}
}
