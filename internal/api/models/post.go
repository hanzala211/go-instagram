package models

import "time"

type Post struct {
	ID        string    `pg:"uuid,pk" json:"id,omitempty"`
	Content   string    `pg:"content" json:"content"`
	Title     string    `pg:"title" json:"title"`
	UserID    string    `pg:"user_id,fk" json:"user_id"`
	User      *User     `pg:"rel:has-one" json:"user"`
	CreatedAt time.Time `json:"createdAt" pg:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" pg:"updatedAt"`
}
