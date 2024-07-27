package main

import (
	"database/sql"
	"fmt"
	"go-myapi/database/models"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 接続に使うユーザー・パスワード・データベース名を定義
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"

	// データベースに接続するためのアドレス文を定義
	// ここでは"docker:docker@tcp(127.0.0.1"3306)/sampledb?parseTime=true"となる
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser,
		dbPassword, dbDatabase)

	// Open関数を用いてデータベースに接続
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	// プログラムが終了するとき、コネクションが close されるようにする
	defer db.Close()

	// データを挿入する処理
	article := models.Article{
		Title:    "insert test",
		Contents: "Can I insert data correctly?",
		UserName: "saki",
	}
	const sqlStr = `
		insert into articles (title, contents, username, nice, created_at) values (?, ?, ?, 0, now());
	`
	// プレースホルダー?に
	// article.TItle, article.Contens, article.UserNameを埋め込んでクエリを実行
	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 結果を確認
	// result.LastInsertId の実行結果から、記事IDが何番になったのかを調べる
	fmt.Println(result.LastInsertId())
	// result.RowsAffected の実行結果から、クエリの影響範囲の広さを調べる
	fmt.Println(result.RowsAffected())
}

/*
トランザクションを利用する流れ
1. トランザクションを貼る（begin）sql.DB型 の Beginメソッド
2. クエリの実行
3. すべてのクエリ実行が成功した場合には、コミットして結果を確定させる（commit）sql.TxのCommitメソッド
4. 一部のクエリ実行が失敗した場合には、ロールバックして結果を無かったことにする（rollback）sql.TxのRollbackメソッド
*/
