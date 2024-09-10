package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

// テスト全体で共有する sql.DB 型
var testDB *sql.DB
var (
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbDatabase = "mydb"
	dbConn     = fmt.Sprintf("postgres://%s:%s@127.0.0.1:5432/%s?sslmode=disable", dbUser, dbPassword, dbDatabase)
)

func connectDB() error {
	var err error
	testDB, err = sql.Open("postgres", dbConn)
	if err != nil {
		return err
	}
	return nil
}

// 全テスト共通の前処理
func setup() error {
	if err := connectDB(); err != nil {
		return err
	}
	return nil
}

// 前テスト共通の後処理
func teardown() {
	testDB.Close()
}

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		os.Exit(1) // m.RunはFatal系メソッドを持たないのでExitする。
	}

	m.Run() // m.Runはテスト関数全てを実行する。

	teardown()
}
