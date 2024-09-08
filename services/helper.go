package services

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbDatabase = "mydb"
	dbConn     = fmt.Sprintf("postgres://%s:%s@127.0.0.1:5432/%s?sslmode=disable", dbUser, dbPassword, dbDatabase)
)

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
