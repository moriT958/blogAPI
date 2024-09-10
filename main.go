package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/moriT958/go-api/api"
)

var (
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbDatabase = "mydb"
	dbConn     = fmt.Sprintf("postgres://%s:%s@127.0.0.1:5432/%s?sslmode=disable", dbUser, dbPassword, dbDatabase)
)

func main() {
	// DBに接続できるかの確認
	// if err := db.Ping(); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("connect to DB")
	// }

	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}
	defer db.Close()

	r := api.NewRouter(db)

	log.Println("server started at port 8080")
	// http.ListenAndServe(":8080", r)はerrを返す。第二引数はnilだと標準のDefaultServeMuxが設定される。
	log.Fatal(http.ListenAndServe(":8080", r)) // errを表示して終了
}
