package model

import "time"

type ClipTagged struct {
	ClipID    string
	TagID     string
	UserID    string
	CreatedAt *time.Time
	UpdatedAt *time.Time

	User *User
	Clip *Clip
	Tag  *Tag
}

func (ClipTagged) TableName() string {
	return "clip_tagged"
}

func NewClipTagged(clipID, tagID, userID string) *ClipTagged {
	return &ClipTagged{
		ClipID:    clipID,
		TagID:     tagID,
		UserID:    userID,
		CreatedAt: nil,
		UpdatedAt: nil,
	}
}
