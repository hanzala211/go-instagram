package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-pg/pg/v10"
	"github.com/go-playground/validator/v10"
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
		if errs, ok := err.(validator.ValidationErrors); ok {
			messages := []string{}
			for _, e := range errs {
				messages = append(messages, fmt.Sprintf("Field '%s' has the issue '%s'", e.Field(), e.Tag()))
			}

			utils.WriteError(w, http.StatusBadRequest, messages)
			return
		}
	}

	post := &models.Post{
		UserID:  user.ID,
		Title:   signupReq.Title,
		Content: signupReq.Content,
	}

	err = h.postService.CreatePost(post)
	fmt.Println(err)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Error Occured!")
		return
	}

	utils.WriteResponse(w, 201, post)
}

func (h *PostHandler) GetPostById(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "postID")
	post := &models.Post{
		ID: postId,
	}
	err := h.postService.GetPostById(post)
	fmt.Println(err)
	if err != nil {
		if err == pg.ErrNoRows {
			utils.WriteError(w, 400, "Not Found")
			return
		}
		utils.WriteError(w, 500, "Error Occured")
		return
	}

	utils.WriteResponse(w, 200, post)
}
