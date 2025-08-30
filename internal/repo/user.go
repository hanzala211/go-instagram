package repo

import (
	"github.com/go-pg/pg/v10"
	"github.com/hanzala211/instagram/internal/api/models"
)

type UserRepo struct {
	db *pg.DB
}

func NewUserRepo(db *pg.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) CreateUser(user *models.User) error {
	_, err := r.db.Model(user).Insert()
	return err
}

func (r *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := r.db.Model(user).Where("email = ?", email).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *UserRepo) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	err := r.db.Model(user).Where("username = ?", username).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *UserRepo) GetUserById(id string) (*models.User, error) {
	user := &models.User{}
	err := r.db.Model(user).Where("id = ?", id).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

