package repositories

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/moriT958/go-api/models"
)

func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
		insert into comments (article_id, message, created_at) values 
		($1, $2, now())
		returning comment_id;
	`
	var newComment models.Comment
	newComment.ArticleID, newComment.Message = comment.ArticleID, comment.Message
	row := db.QueryRow(sqlStr, newComment.ArticleID, newComment.Message)
	if err := row.Scan(&newComment.CommentID); err != nil {
		return models.Comment{}, err
	}

	return newComment, nil
}

func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `select * from comments where article_id = $1;`

	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	commentArray := make([]models.Comment, 0)
	for rows.Next() {
		var comment models.Comment
		var createdTime sql.NullTime

		if err := rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &createdTime); err != nil {
			return nil, err
		}

		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}

		commentArray = append(commentArray, comment)
	}

	return commentArray, nil
}
