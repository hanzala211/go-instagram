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

// func (r *PostRepo) GetPostSuggestions(userId string) ([]*models.Post, error) {
// 	var posts []*models.Post
// 	err := r.db.Model(&posts).
// 		ColumnExpr("post.*, (SELECT COUNT(*) FROM likes WHERE likes.post_id = post.id) AS likes_count").
// 		Join("LEFT JOIN posts AS liked_posts ON liked_posts.user_id = post.user_id").
// 		Join("LEFT JOIN likes l ON l.post_id = liked_posts.id").
// 		Join("LEFT JOIN likes l2 ON l2.post_id = post.id AND l2.user_id = ?", userId).
// 		Where("l2.post_id IS NULL").
// 		OrderExpr("likes_count DESC, CASE WHEN l.user_id IS NOT NULL THEN 0 ELSE 1 END, RANDOM()").Select()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return posts, nil
// }

func (r *PostRepo) GetPostSuggestions(userId string) ([]*models.Post, error) {
	var posts []*models.Post
	err := r.db.Model(&posts).ColumnExpr(`post.*, 0.5 * (SELECT COUNT(*) FROM likes WHERE likes.post_id = post.id) +
		0.3 * (-EXTRACT(EPOCH FROM NOW() - post."createdAt") / 86400) +
		1.0 *  CASE WHEN post.user_id IN (SELECT liked_posts.user_id FROM likes
			JOIN posts AS liked_posts ON liked_posts.id = likes.post_id
			WHERE likes.user_id = ?
		) THEN 1 ELSE 0 END AS _score`,
		userId).
		Relation("User").Relation("Likes").
		OrderExpr("_score DESC").
		Select()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

// func (r *PostRepo) GetPostSuggestions(userId string) ([]*models.Post, error) {
// 	var posts []*models.Post
// 	err := r.db.Model(&posts).ColumnExpr(`post.*,
// 		COUNT(*) AS likes_count,
// 		0.5 * COUNT(*) +
// 		0.3 * (-EXTRACT(EPOCH FROM NOW() - post."createdAt") / 86400) +
// 		1.0 * CASE WHEN affinity_users.user_id IS NOT NULL THEN 1 ELSE 0 END AS _score
// 		`).
// 		Join("LEFT JOIN likes ON likes.post_id = post.id").
// 		Join(`LEFT JOIN
// 		(
// 		 SELECT DISTINCT liked_posts.user_id
// 		 FROM likes
// 		 JOIN posts AS liked_posts ON liked_posts.id = likes.post_id
// 		 WHERE likes.user_id = ?
// 		) AS affinity_users ON affinity_users.user_id = post.user_id
// 		`, userId).
// 		GroupExpr("post.id, affinity_users.user_id").
// 		OrderExpr("_score DESC").
// 		Select()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return posts, nil
// }
