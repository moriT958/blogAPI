package repositories

import (
	"database/sql"

	"github.com/moriT958/go-api/models"
)

type ICommentRepositoy interface {
	Create(models.Comment) (models.Comment, error)
	FindAll(int) ([]models.Comment, error)
}

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (r *CommentRepository) Create(comment models.Comment) (models.Comment, error) {
	query := "INSERT INTO comments (article_id, message, created_at) VALUES (?, ?, now());"

	res, err := r.db.Exec(query, comment.ArticleID, comment.Message)
	if err != nil {
		return models.Comment{}, err
	}

	commentID, err := res.LastInsertId()
	if err != nil {
		return models.Comment{}, err
	}
	comment.CommentID = int(commentID)

	return comment, nil
}

func (r *CommentRepository) FindAll(articleID int) ([]models.Comment, error) {
	rows, err := r.db.Query("SELECT * FROM comments WHERE article_id = ?;", articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comment models.Comment
	comments := make([]models.Comment, 0)
	for rows.Next() {
		var createdTime sql.NullTime
		if err := rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &createdTime); err != nil {
			return nil, err
		}

		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}

		comments = append(comments, comment)
	}

	return comments, nil
}
