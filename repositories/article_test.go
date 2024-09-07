package repositories_test

import (
	"testing"

	_ "github.com/lib/pq"
	"github.com/moriT958/go-api/models"
	"github.com/moriT958/go-api/repositories"
	"github.com/moriT958/go-api/repositories/testdata"
)

// SelectArticleList関数のテスト
func TestSelectArticleList(t *testing.T) {
	expectedNum := len(testdata.ArticleTestData)
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}
}

func TestSelectArticleDetail(t *testing.T) {
	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "subtest1",
			expected:  testdata.ArticleTestData[0],
		},
		{
			testTitle: "subtest2",
			expected:  testdata.ArticleTestData[1],
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				t.Fatal(err) // FatalとErrorの違いは処理が落ちるか続くか。
			}

			if got.ID != test.expected.ID {
				t.Errorf("ID: get %d but want %d\n", got.ID, test.expected.ID)
			}
			if got.Title != test.expected.Title {
				t.Errorf("Title: get %s but want %s\n", got.Title, test.expected.Title)
			}
			if got.Contents != test.expected.Contents {
				t.Errorf("Content: get %s but want %s\n", got.Contents, test.expected.Contents)
			}
			if got.UserName != test.expected.UserName {
				t.Errorf("UserName: get %s but want %s\n", got.UserName, test.expected.UserName)
			}
			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, test.expected.NiceNum)
			}
		})
	}
}

// InsertArticle関数のテスト
func TestInsertArticle(t *testing.T) {
	testArticle := models.Article{
		Title:    "insertTest",
		Contents: "testest",
		UserName: "moriT",
	}

	expectedArticleNum := 3 // postgresの自動採番はシーケンス方式なのでここの値注意！
	newArticle, err := repositories.InsertArticle(testDB, testArticle)
	if err != nil {
		t.Error(err) // ErrorはFatalと違い、以後の処理を継続。
	}
	if newArticle.ID != expectedArticleNum {
		t.Errorf("new article id is expected %d but got %d\n", expectedArticleNum, newArticle.ID)
	}

	// InsertArticle個別の後処理
	t.Cleanup(func() {
		const sqlStr = `
			delete from articles
			where title = $1 and contents = $2 and username = $3;
		`
		testDB.Exec(sqlStr, testArticle.Title, testArticle.Contents, testArticle.UserName)
	})
}

func TestUpdateNiceNum(t *testing.T) {
	testID := 1

	before, err := repositories.SelectArticleDetail(testDB, testID)
	if err != nil {
		t.Fatal("fail to get before data")
	}

	err = repositories.UpdateNiceNum(testDB, testID)
	if err != nil {
		t.Fatal(err)
	}

	after, err := repositories.SelectArticleDetail(testDB, testID)
	if err != nil {
		t.Fatal("fail to get after data")
	}

	if after.NiceNum-before.NiceNum != 1 {
		t.Error("fail to update nice num")
	}

	// 後処理
	t.Cleanup(func() {
		const sqlStr = `update articles set nice = nice - 1 where article_id = $1;`
		_, err := testDB.Exec(sqlStr, testID)
		if err != nil {
			t.Error("fail to cleanup updateNiceNum")
		}
	})
}
