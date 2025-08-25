package repo

import "github.com/hanzala211/instagram/internal/api/models"

type UserStorage interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserById(id string) (*models.User, error)

}

type Storage struct {
	User UserStorage
}

func NewStorage(userRepo *UserRepo) *Storage {
	return &Storage{
		User: userRepo,
	}
}