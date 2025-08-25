package services

import (
	"errors"

	"github.com/hanzala211/instagram/internal/api/models"
	"github.com/hanzala211/instagram/internal/repo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	store *repo.Storage
}

func NewUserService(store *repo.Storage) *UserService {
	return &UserService{
		store: store,
	}
}

func (s *UserService) CreateUser(user *models.User) error {
	existingUser, err := s.store.User.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("Email already exists")
	}
	existingUser, err = s.store.User.GetUserByUsername(user.Username)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("Username already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.store.User.CreateUser(user)
}

func (s *UserService) Login(user *models.User) (existingUser *models.User, err error) {
	existingUser, err = s.store.User.GetUserByUsername(user.Username)
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, errors.New("User not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		return nil, errors.New("Invalid password")
	}
	return existingUser, nil
}

func (s *UserService) GetUserById(id string) (existingUser *models.User, err error) {
	existingUser, err = s.store.User.GetUserById(id)
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, errors.New("User not found")
	}
	return existingUser, nil
}