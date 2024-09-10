package controllers_test

import (
	"testing"

	"github.com/moriT958/go-api/controllers"
	"github.com/moriT958/go-api/controllers/testdata"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)

	m.Run()
}
