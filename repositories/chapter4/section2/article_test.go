/*
テーブルドリブンテストは主に以下のような流れになる
1.「テストケース名」と「テストデータ」セットのスライスを作成
2.1で作ったものをfor文で回す
3.2の中でサブテストを実施
*/

package repositories_test

import (
	"database/sql"
	"fmt"
	"go-myapi/models"
	"go-myapi/repositories"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArticleDetail(t *testing.T) {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	// データベースに接続する
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		// 接続できなかった場合にはテストそのものが続行不可
		// -> Fatal で終了させる
		t.Fatal(err)
	}
	defer db.Close()

	// 1.「テストケース名」と「テストデータ」セットのスライスを作成
	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "subtest1",
			expected: models.Article{
				ID:       1,
				Title:    "firstPost",
				Contents: "This is my first blog",
				UserName: "saki",
				NiceNum:  3,
			},
		}, {
			testTitle: "subtest2",
			expected: models.Article{
				ID:       2,
				Title:    "firstPost",
				Contents: "This is my first blog",
				UserName: "saki",
				NiceNum:  2,
			},
		},
	}

	// 2. 1で作ったスライスをfor文で回す
	for _, test := range tests {
		// 3. サブテストを実施
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(db, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}
			if got.ID != test.expected.ID {
				t.Errorf("ID: get %d but want %d\n", got.ID, test.expected.ID)
			}
			if got.Title != test.expected.Title {
				t.Errorf("Title: get %s but want %s\n", got.Title, test.expected.Title)
			}
			if got.Contents != test.expected.Contents {
				t.Errorf("Contents: get %s but want %s\n", got.Contents, test.expected.Contents)
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
