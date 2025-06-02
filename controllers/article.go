package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/moriT958/go-api/models"
)

const jsonContentType = "application/json"

func (c *Controller) PostArticleHandler(w http.ResponseWriter, req *http.Request) {

	var reqArticle reqPostArticle
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		log.Println(err)
		http.Error(w, "入力値が不正です。", http.StatusBadRequest)
		return
	}

	newArticle := models.Article{
		Title:    reqArticle.Title,
		Contents: reqArticle.Contents,
		UserName: reqArticle.UserName,
	}

	a, err := c.Service.PostArticle(newArticle)
	if err != nil {
		log.Println(err)
		http.Error(w, "記事の投稿に失敗しました。", http.StatusInternalServerError)
		return
	}

	resp := respPostArtile{
		ID: a.ID,
	}
	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println(err)
		http.Error(w, "データの整形に失敗しました。", http.StatusInternalServerError)
		return
	}
}

func (c *Controller) GetArticlesHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	articles, err := c.Service.GetArticles(page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := new(respGetArticles)
	resp.Articles = make([]baseArticle, len(articles))
	for i, a := range articles {
		resp.Articles[i] = baseArticle{
			ID:        a.ID,
			Title:     a.Title,
			Contents:  a.Contents,
			UserName:  a.UserName,
			NiceNum:   a.NiceNum,
			CreatedAt: a.CreatedAt,
		}
	}
	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (c *Controller) GetArticleDetailHandler(w http.ResponseWriter, req *http.Request) {

	idStr := req.PathValue(IdPathVal)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "idは整数値で入力してください。", http.StatusBadRequest)
		return
	}

	article, err := c.Service.GetArticle(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "記事の取得に失敗しました。", http.StatusInternalServerError)
		return
	}

	resp := &respGetArticleDetail{
		ID:        article.ID,
		Title:     article.Title,
		Contents:  article.Contents,
		UserName:  article.UserName,
		NiceNum:   article.NiceNum,
		Comments:  make([]baseComment, len(article.CommentList)),
		CreatedAt: article.CreatedAt,
	}
	for i, c := range article.CommentList {
		resp.Comments[i] = baseComment{
			ID:        c.CommentID,
			Message:   c.Message,
			CreatedAt: c.CreatedAt,
		}
	}
	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (c *Controller) PatchNiceHandler(w http.ResponseWriter, req *http.Request) {
	idStr := req.PathValue(IdPathVal)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "idは整数値で入力してください。", http.StatusBadRequest)
		return
	}

	article, err := c.Service.AddNice(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "いいねの更新に失敗しました。", http.StatusInternalServerError)
		return
	}

	resp := baseArticle{
		ID:        article.ID,
		Title:     article.Title,
		Contents:  article.Contents,
		UserName:  article.UserName,
		NiceNum:   article.NiceNum,
		CreatedAt: article.CreatedAt,
	}
	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
