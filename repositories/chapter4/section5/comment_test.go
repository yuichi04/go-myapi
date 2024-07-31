package repositories_test

import (
	"go-myapi/repositories"
	"go-myapi/repositories/testdata"
	"testing"
)

// SelectCommentList
// 指定した記事のコメントリストが取得されること
func TestSelectCommentList(t *testing.T) {
	got, err := repositories.SelectCommentList(testDB, testdata.ArticleTestData[0].ID)
	if err != nil {
		t.Fatal(err)
	}

	for _, comment := range got {
		if comment.ArticleID != testdata.ArticleTestData[0].ID {
			t.Errorf("want comment of articleID %d but got ID %d\n",
				comment.ArticleID,
				testdata.ArticleTestData[0].ID,
			)
		}
	}
}

func TestInsertComment(t *testing.T) {
	newComment, err := repositories.InsertComment(testDB, testdata.InsertComment)
	if err != nil {
		t.Fatal(err)
	}
	if newComment.Message != testdata.InsertComment.Message {
		t.Errorf("want new comment message %s but got message %s\n",
			newComment.Message,
			testdata.InsertComment.Message,
		)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from comments
			where message = ?
		`
		testDB.Exec(sqlStr, testdata.InsertComment.Message)
	})
}
