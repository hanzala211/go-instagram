package services

import (
	"github.com/hanzala211/instagram/internal/api/models"
	"github.com/hanzala211/instagram/internal/repo"
)

type PostService struct {
	store *repo.Storage
}

func NewPostService(store *repo.Storage) *PostService {
	return &PostService{
		store: store,
	}
}

func (s *PostService) CreatePost(post *models.Post) error {
	err := s.store.Post.CreatePost(post)
	return err
}
