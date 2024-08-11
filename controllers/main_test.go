package controllers_test

import (
	"go-myapi/controllers"
	"go-myapi/controllers/testdata"
	"testing"
)

// 1. テストに使うリソース（コントローラ構造体）を用意
var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)

	m.Run()
}

// ミドルウェアとはアプリケーションがそのビジネスロジック部分を実行する前後に行う追加の処理を実行する部分のこと
