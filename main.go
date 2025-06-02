package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/moriT958/go-api/api"
)

func main() {

	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")
	host := os.Getenv("MYSQL_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", username, password, host, database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("fail to connect DB:", err)
		return
	}
	defer db.Close()

	r := api.NewRouter(db)

	log.Println("server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}
