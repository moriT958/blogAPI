package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/moriT958/go-api/api"
	//"github.com/moriT958/go-api/models"
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

	// トランザクションの開始
	// tx, err := db.Begin()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// // 現在のいいね数を取得するクエリを実行する
	// article_id := 1
	// const sqlGetNice = `
	// 	select nice
	// 	from articles
	// 	where article_id = $1;
	// `

	// row := tx.QueryRow(sqlGetNice, article_id) // QueryもExecもtxのインターフェースとして用意されているので、dbと同様に使用可能。
	// if err := row.Err(); err != nil {
	// 	fmt.Println(err)
	// 	tx.Rollback() // 失敗したらロールバック
	// 	return
	// }

	// // 変数 nicenum に現在のいいね数を読み込む
	// var nicenum int
	// err = row.Scan(&nicenum)
	// if err != nil {
	// 	fmt.Println(err)
	// 	tx.Rollback()
	// 	return
	// }
	// // いいね数を+1 する更新処理を行う
	// const sqlUpdateNice = `update articles set nice = $1 where article_id = $2`
	// _, err = tx.Exec(sqlUpdateNice, nicenum+1, article_id)
	// if err != nil {
	// 	fmt.Println(err)
	// 	tx.Rollback()
	// 	return
	// }
	// // コミットして処理内容を確定させる
	// tx.Commit()

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
