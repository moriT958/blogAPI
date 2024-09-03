package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moriT958/go-api/handlers"
)

func main() {

	r := mux.NewRouter()

	// muxを使うことで,ルーティング時にメソッドを絞ることができる。
	// mux側が内部で自動にhttp.Errorを返す。
	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server started at port 8080")
	err := http.ListenAndServe(":8080", r) // 第二引数はnilだと標準のDefaultServeMuxが設定される。
	log.Fatal(err)                         // errを表示して終了
}
