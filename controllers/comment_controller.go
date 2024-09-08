package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/moriT958/go-api/controllers/services"
	"github.com/moriT958/go-api/models"
)

type CommentController struct {
	service services.CommentServicer
}

func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

// POST /comment のハンドラ
func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	// comment1 := models.Comment1
	// jsonData, err := json.Marshal(comment1)
	// if err != nil {
	// 	http.Error(w, "Faild to encode json", http.StatusInternalServerError)
	// 	return
	// }

	// w.Write(jsonData)

	// 上記のリファクタ
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to encode json\n", http.StatusBadRequest)
	}

	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	resComment := comment
	if err := json.NewEncoder(w).Encode(resComment); err != nil {
		http.Error(w, "fail to encode json\n", http.StatusBadRequest)
	}
}
