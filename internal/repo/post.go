package repo

import (
	"github.com/go-pg/pg/v10"
	"github.com/hanzala211/instagram/internal/api/models"
)

type PostRepo struct {
	db *pg.DB
}

func NewPostRepo(db *pg.DB) *PostRepo {
	return &PostRepo{
		db: db,
	}
}

func (r *PostRepo) CreatePost(post *models.Post) error {
	_, err := r.db.Model(post).Insert()
	return err
}

func (r *PostRepo) GetPostById(post *models.Post) error {
	err := r.db.Model(post).WherePK().Relation("User").Select()
	return err
}
