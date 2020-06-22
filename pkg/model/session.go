package model

import "time"

type Session struct {
	ID        string
	UserID    string
	ExpireAt  *time.Time
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func NewSession(userID string, expireAt *time.Time) *Session {
	return &Session{
		ID:        GetIDGenerator().New(),
		UserID:    userID,
		ExpireAt:  expireAt,
		CreatedAt: nil,
		UpdatedAt: nil,
	}
}
