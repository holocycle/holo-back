package model

import "time"

type ClipTag struct {
	ID        string
	UserID    string
	ClipID    string
	TagID     string
	CreatedAt *time.Time
	UpdatedAt *time.Time

	User *User
	Clip *Clip
	Tag  *Tag
}

func NewClipTag(userID, clipID, tagID string) *ClipTag {
	return &ClipTag{
		ID:        NewID(),
		UserID:    userID,
		ClipID:    clipID,
		TagID:     tagID,
		CreatedAt: nil,
		UpdatedAt: nil,
	}
}
