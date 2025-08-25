package models

import "time"

type User struct {
	ID        string    `json:"id,omitempty" pg:"type:uuid,pk"`
	FullName  string    `json:"fullName" pg:"fullName"`
	Username  string    `json:"username" pg:"username"`
	Email     string    `json:"email" pg:"email"`
	Password  string    `json:"-" pg:"password"`
	CreatedAt time.Time `json:"createdAt" pg:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" pg:"updatedAt"`
}