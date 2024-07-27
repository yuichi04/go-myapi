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

	// 1. テスト結果として期待する値を定義
	expected := models.Article{
		ID:       1,
		Title:    "firstPost",
		Contents: "This is my first blog",
		UserName: "saki",
		NiceNum:  3,
	}

	// 2.テスト対象となる関数を実行
	got, err := repositories.SelectArticleDetail(db, expected.ID)
	if err != nil {
		// SelectArticleDetailがうまくいかなくて、そもそも戻り値gotが得られていないなら
		// この後の期待する出力と実際の出力の比較が不可能
		// -> テスト続行不可なのでFatalで終了させる
		t.Fatal(err)
	}

	// 3. 2の結果と1の値を比べる
	if got.ID != expected.ID {
		t.Errorf("ID: get %d but want %d\n", got.ID, expected.ID)
	}
	if got.Title != expected.Title {
		t.Errorf("Title: get %s but want %s\n", got.Title, expected.Title)
	}
	if got.Contents != expected.Contents {
		t.Errorf("Contents: get %s but want %s\n", got.Contents, expected.Contents)
	}
	if got.UserName != expected.UserName {
		t.Errorf("UserName: get %s but want %s\n", got.UserName, expected.UserName)
	}
	if got.NiceNum != expected.NiceNum {
		t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, expected.NiceNum)
	}
}
