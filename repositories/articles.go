package repositories

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/moriT958/go-api/models"
)

const (
	articleNumPerPage = 5
)

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
		INSERT INTO articles (title, contents, username, nice, created_at) VALUES
		($1, $2, $3, 0, now())
		RETURNING article_id
		;
	`
	var newArticle models.Article
	newArticle.Title, newArticle.Contents, newArticle.UserName = article.Title, article.Contents, article.UserName

	row := db.QueryRow(sqlStr, newArticle.Title, newArticle.Contents, newArticle.UserName)
	if err := row.Scan(&newArticle.ID); err != nil {
		return models.Article{}, err
	}

	return newArticle, nil
}

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
		select article_id, title, contents, username, nice from articles
		limit $1 
		offset $2;
	`

	rows, err := db.Query(sqlStr, articleNumPerPage, articleNumPerPage*(page-1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articleArray := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article
		rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)
		articleArray = append(articleArray, article)
	}

	return articleArray, nil
}

func SelectArticleDetail(db *sql.DB, id int) (models.Article, error) {
	const sqlStr = `
		select * from articles where article_id = $1;
	`
	row := db.QueryRow(sqlStr, id)
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

func UpdateNiceNum(db *sql.DB, id int) error {
	const sqlStr = `
		update articles set nice = nice + 1 where article_id = $1;
	`
	_, err := db.Exec(sqlStr, id)
	if err != nil {
		return err
	}

	return nil
}
