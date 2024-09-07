package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/moriT958/go-api/handlers"
	//"github.com/moriT958/go-api/models"
)

func main() {
	dbUser := "postgres"
	dbPassword := "postgres"
	dbDatabase := "mydb"
	dbConn := fmt.Sprintf("postgres://%s:%s@127.0.0.1:5432/%s?sslmode=disable", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// DBに接続できるかの確認
	// if err := db.Ping(); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("connect to DB")
	// }

	// クエリの定義
	// articleID := 1
	// const sqlStr = `
	//     select *
	//     from articles
	//     where article_id = $1;
	// `

	// クエリの実行
	// Queryはレコードを返り値として返す。
	// rows, err := db.Query(sqlStr, articleID) // クエリの埋め込みはSQLインジェクション対策が施されてるQueryを使う。fmtは使わない。
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer rows.Close()

	// 返ってくるデータが0か1件と確定な場合はQueryRowを使う。
	// row := db.QueryRow(sqlStr, articleID)
	// if err := row.Err(); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// var article models.Article
	// var createdTime sql.NullTime // created_atフィールドがNULLの可能性があるので

	// articleArray := make([]models.Article, 0)
	// for rows.Next() {
	// 	err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)

	// 	if createdTime.Valid { // NULLじゃなかったらValidはtrueとなる.
	// 		article.CreatedAt = createdTime.Time
	// 	}

	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		articleArray = append(articleArray, article)
	// 	}
	// }
	// fmt.Printf("%+v\n", articleArray)

	// err = row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// if createdTime.Valid {
	// 	article.CreatedAt = createdTime.Time
	// }
	// fmt.Printf("%+v\n", article)

	// データを挿入する処理
	// article := models.Article{
	// 	Title:    "insert test",
	// 	Contents: "Can I insert data correctly?",
	// 	UserName: "moriT",
	// }

	// const sqlStr = `
	// 	insert into articles (title, contents, username, nice, created_at) values
	// 	($1, $2, $3, 0, now())
	// 	;
	// `
	// // レコードを返して欲しい時はQuery、resultを返して欲しい時(insert)はExecを使う。
	// result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// 結果を確認
	// fmt.Println(result.LastInsertId()) // 何行目に追加されたかを返す。(psqlのドライバは未対応)
	// fmt.Println(result.RowsAffected()) // 何行追加されたかを返す。

	// トランザクションの開始
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 現在のいいね数を取得するクエリを実行する
	article_id := 1
	const sqlGetNice = `
		select nice
		from articles
		where article_id = $1;
	`

	row := tx.QueryRow(sqlGetNice, article_id) // QueryもExecもtxのインターフェースとして用意されているので、dbと同様に使用可能。
	if err := row.Err(); err != nil {
		fmt.Println(err)
		tx.Rollback() // 失敗したらロールバック
		return
	}

	// 変数 nicenum に現在のいいね数を読み込む
	var nicenum int
	err = row.Scan(&nicenum)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}
	// いいね数を+1 する更新処理を行う
	const sqlUpdateNice = `update articles set nice = $1 where article_id = $2`
	_, err = tx.Exec(sqlUpdateNice, nicenum+1, article_id)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}
	// コミットして処理内容を確定させる
	tx.Commit()

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
	// http.ListenAndServe(":8080", r)はerrを返す。第二引数はnilだと標準のDefaultServeMuxが設定される。
	log.Fatal(http.ListenAndServe(":8080", r)) // errを表示して終了
}
