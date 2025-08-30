package dto

type SignupRequest struct {
	Username string `json:"username" validate:"required,alphanum,min=8,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
	FullName string `json:"fullName" validate:"required,min=3,max=100"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,alphanum,min=8,max=32"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

