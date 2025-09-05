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

func (s *PostService) GetPostById(post *models.Post) error {
	err := s.store.Post.GetPostById(post)
	return err
}

func (s *PostService) GetPostsForUser(userId string) ([]*models.Post, error) {
	return s.store.Post.GetUserPosts(userId)
}

func (s *PostService) LikePost(postId string, userId string) error {
	like := &models.Likes{
		PostID: postId,
		UserID: userId,
	}
	err := s.store.Post.LikePost(like)
	return err
}

func (s *PostService) DislikePost(postId string, userId string) error {
	err := s.store.Post.DislikePost(postId, userId)
	return err
}
