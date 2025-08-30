package handler

import (
	"encoding/json"
	"net/http"

	"github.com/hanzala211/instagram/internal/api/dto"
	"github.com/hanzala211/instagram/internal/api/models"
	"github.com/hanzala211/instagram/internal/services"
	"github.com/hanzala211/instagram/utils"
)

type PostHandler struct {
	postService *services.PostService
}

func NewPostHandler(postService *services.PostService) *PostHandler {
	return &PostHandler{
		postService: postService,
	}
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*models.User)
	signupReq := &dto.CreatePostReq{}
	json.NewDecoder(r.Body).Decode(signupReq)
	err := validate.Struct(signupReq)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	post := &models.Post{
		UserID:  user.ID,
		Title:   signupReq.Title,
		Content: signupReq.Content,
	}

	err = h.postService.CreatePost(post)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Error Occured!")
		return
	}

	utils.WriteResponse(w, 201, post)
}
