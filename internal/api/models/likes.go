package models

import "time"

type Likes struct {
	UserID    string    `json:"user_id" pg:"type:uuid,pk"`
	PostID    string    `json:"post_id" pg:"type:uuid,pk"`
	User      *User     `json:"user,omitempty" pg:"rel:has-one"`
	Post      *Post     `json:"post,omitempty" pg:"rel:has-one"`
	CreatedAt time.Time `json:"createdAt" pg:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" pg:"updatedAt"`
}
