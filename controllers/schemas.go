package controllers

import "time"

/* Base Schema */

type baseArticle struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Contents  string    `json:"contents"`
	UserName  string    `json:"userName"`
	NiceNum   int       `json:"nice"`
	CreatedAt time.Time `json:"createdAt"`
}

type baseComment struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}

/* Json Schemas */

type reqPostArticle struct {
	Title    string `json:"title"`
	Contents string `json:"contents"`
	UserName string `json:"userName"`
}

type respPostArtile struct {
	ID int `json:"id"`
}

type respGetArticles struct {
	Articles []baseArticle `json:"articles"`
}

type respGetArticleDetail struct {
	ID        int           `json:"id"`
	Title     string        `json:"title"`
	Contents  string        `json:"contents"`
	UserName  string        `json:"userName"`
	NiceNum   int           `json:"nice"`
	Comments  []baseComment `json:"comments"`
	CreatedAt time.Time     `json:"createdAt"`
}

type reqPostComment struct {
	ArticleID int    `json:"articleId"`
	Message   string `json:"message"`
}

type respPostComment struct {
	ID        int    `json:"id"`
	ArticleID int    `json:"articleId"`
	Message   string `json:"message"`
}

/* Path Values */

const IdPathVal = "id"
