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
	err := r.db.Model(post).
		WherePK().
		Relation("User").
		Relation("Likes").
		Join("LEFT JOIN likes ON likes.post_id = post.id").
		ColumnExpr("post.*, COUNT(likes.post_id) AS likes_count").
		Group("post.id", "user.id").
		Select()
	return err
}

func (r *PostRepo) GetUserPosts(userId string) ([]*models.Post, error) {
	var posts []*models.Post
	err := r.db.Model(&posts).Where("user_id = ?", userId).
		ColumnExpr("*, (SELECT COUNT(*) FROM likes WHERE likes.post_id = post.id) as likes_count").
		Select()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostRepo) LikePost(like *models.Likes) error {
	_, err := r.db.Model(like).Insert()
	return err
}

func (r *PostRepo) DislikePost(postId string, userId string) error {
	_, err := r.db.Model(&models.Likes{PostID: postId, UserID: userId}).WherePK().Delete()
	return err
}
