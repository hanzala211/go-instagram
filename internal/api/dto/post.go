package dto

type CreatePostReq struct {
	Title   string `json:"title" validate:"required,min=3,max=32"`
	Content string `json:"content" validate:"required,min=3,max=150"`
}
