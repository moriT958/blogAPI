package repositories

import (
	"database/sql"

	"github.com/moriT958/go-api/models"
)

type IArticleRepository interface {
	Create(models.Article) (models.Article, error)
	FindAll(int) ([]models.Article, error)
	FindByID(int) (models.Article, error)
	AddNice(int) error
}

var _ IArticleRepository = (*ArticleRepository)(nil)

type ArticleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) *ArticleRepository {
	return &ArticleRepository{
		db: db,
	}
}

func (r *ArticleRepository) Create(article models.Article) (models.Article, error) {
	query := "INSERT INTO articles (title, contents, username, nice, created_at) VALUES (?, ?, ?, 0, now());"

	res, err := r.db.Exec(query, article.Title, article.Contents, article.UserName)
	if err != nil {
		return models.Article{}, err
	}

	articleID, err := res.LastInsertId()
	if err != nil {
		return models.Article{}, err
	}
	article.ID = int(articleID)

	return article, nil
}

const articleNumPerPage = 5

func (r *ArticleRepository) FindAll(page int) ([]models.Article, error) {
	query := "SELECT id, title, contents, username, nice FROM articles LIMIT ? OFFSET ?;"

	rows, err := r.db.Query(query, articleNumPerPage, articleNumPerPage*(page-1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var article models.Article
	articles := make([]models.Article, 0)
	for rows.Next() {
		rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)
		articles = append(articles, article)
	}

	return articles, nil
}

func (r *ArticleRepository) FindByID(id int) (models.Article, error) {
	query := "SELECT id, title, contents, username, nice, created_at FROM articles WHERE id = ?;"
	row := r.db.QueryRow(query, id)
	if err := row.Err(); err != nil {
		return models.Article{}, err
	}

	var article models.Article
	var createdTime sql.NullTime
	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		return models.Article{}, err
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return article, nil
}

func (r *ArticleRepository) AddNice(id int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	row := tx.QueryRow("SELECT nice FROM articles WHERE id = ?;", id)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return err
	}

	var niceNum int
	err = row.Scan(&niceNum)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("UPDATE articles SET nice = ? WHERE id = ?;", niceNum+1, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
