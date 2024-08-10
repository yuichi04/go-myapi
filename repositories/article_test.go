package repositories_test

import (
	"go-myapi/models"
	"go-myapi/repositories"
	"go-myapi/repositories/testdata"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// SelectArticleList
// 取得された記事数と期待する記事数が一致すること
func TestSelectArticleList(t *testing.T) {
	expectedNum := len(testdata.ArticleTestData)
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d got %d articles\n", expectedNum, num)
	}
}

// SelectArticleDetail
// 取得した記事の詳細と期待する記事の詳細が一致すること
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
				t.Fatal(err)
			}
			if got.ID != test.expected.ID {
				t.Errorf("ID: get %d but want %d\n", got.ID, test.expected.ID)
			}
		})
	}
}

// InsertArticle
// インサートした記事データと返された記事データが同一であること
func TestInsertArticle(t *testing.T) {
	newArticle, err := repositories.InsertArticle(testDB, testdata.InsertArticleData[0])
	if err != nil {
		t.Error(err)
	}
	if newArticle.Title != testdata.InsertArticleData[0].Title {
		t.Errorf("new article title is expected %s but got %s\n",
			testdata.InsertArticleData[0].Title,
			newArticle.Title,
		)
	}
	if newArticle.Contents != testdata.InsertArticleData[0].Contents {
		t.Errorf("new article contents is expected %s but got %s\n",
			testdata.InsertArticleData[0].Contents,
			newArticle.Contents,
		)
	}
	if newArticle.UserName != testdata.InsertArticleData[0].UserName {
		t.Errorf("new article username is expected %s but got %s\n",
			testdata.InsertArticleData[0].UserName,
			newArticle.UserName,
		)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from articles
			where title = ? and contents = ? and username = ?
		`
		testDB.Exec(sqlStr,
			testdata.InsertArticleData[0].Title,
			testdata.InsertArticleData[0].Contents,
			testdata.InsertArticleData[0].UserName,
		)
	})
}

// UpdateNiceNum
// 指定した記事のいいね数が1増加すること
func TestUpdateNiceNum(t *testing.T) {
	before, err := repositories.SelectArticleDetail(testDB, testdata.ArticleTestData[0].ID)
	if err != nil {
		t.Fatal("fail to get before data")
	}

	err = repositories.UpdateNiceNum(testDB, testdata.ArticleTestData[0].ID)
	if err != nil {
		t.Fatal(err)
	}

	after, err := repositories.SelectArticleDetail(testDB, testdata.ArticleTestData[0].ID)
	if err != nil {
		t.Fatal("fail to get after data")
	}

	if after.NiceNum-before.NiceNum != 1 {
		t.Error("fail to update nice num")
	}

	t.Cleanup(func() {
		repositories.DecreaseNiceNum(testDB, testdata.ArticleTestData[0].ID)
	})
}
