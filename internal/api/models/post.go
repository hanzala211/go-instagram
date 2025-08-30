package models

import "time"

type Post struct {
	ID        string    `json:"id,omitempty" pg:"type:uuid,pk"`
	Content   string    `pg:"content" json:"content"`
	Title     string    `pg:"title" json:"title"`
	UserID    string    `pg:"type:uuid,notnull" json:"user_id"`
	User      *User     `pg:"rel:has-one" json:"user"`
	CreatedAt time.Time `json:"createdAt" pg:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" pg:"updatedAt"`
}
