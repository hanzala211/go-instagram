package repo

import "github.com/hanzala211/instagram/internal/api/models"

type UserStorage interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserById(id string) (*models.User, error)
}

type PostStorage interface {
	CreatePost(post *models.Post) error
}

type Storage struct {
	User UserStorage
	Post PostStorage
}

func NewStorage(userRepo *UserRepo, postRepo *PostRepo) *Storage {
	return &Storage{
		User: userRepo,
		Post: postRepo,
	}
}
