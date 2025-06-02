package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/moriT958/go-api/models"
)

func (c *Controller) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment reqPostComment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		log.Println(err)
		http.Error(w, "リクエストデータが不正です。", http.StatusBadRequest)
		return
	}

	newComment := models.Comment{
		ArticleID: reqComment.ArticleID,
		Message:   reqComment.Message,
	}

	comment, err := c.Service.PostComment(newComment)
	if err != nil {
		log.Println(err)
		http.Error(w, "コメントの投稿に失敗しました。", http.StatusInternalServerError)
		return
	}

	resp := respPostComment{
		ID:        comment.CommentID,
		ArticleID: comment.ArticleID,
		Message:   comment.Message,
	}
	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
