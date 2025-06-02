package api

import (
	"database/sql"
	"net/http"

	"github.com/moriT958/go-api/controllers"
	"github.com/moriT958/go-api/repositories"
	"github.com/moriT958/go-api/services"
)

func NewRouter(db *sql.DB) *http.ServeMux {
	ar := repositories.NewArticleRepository(db)
	cr := repositories.NewCommentRepository(db)
	s := services.NewBlogService(ar, cr)
	c := controllers.NewController(s)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /article", c.PostArticleHandler)
	mux.HandleFunc("GET /article/list", c.GetArticlesHandler)
	mux.HandleFunc("GET /article/{id}", c.GetArticleDetailHandler)
	mux.HandleFunc("PATCH /article/{id}", c.PatchNiceHandler)
	mux.HandleFunc("POST /comment", c.PostCommentHandler)

	return mux
}
