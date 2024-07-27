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

	articleID := 1000
	const sqlStr = `
		select *
		from articles
		where article_id = ?;
	`

	// sqlStr内に埋め込まれたプレースホルダー?に
	// articleID の値を入れてクエリを実行
	row := db.QueryRow(sqlStr, articleID)
	// Err メソッドの中身を確認
	if err := row.Err(); err != nil {
		// データ取得件数が0件だった場合は
		// データ読み出し処理には移らずに終了
		fmt.Println(err)
		return
	}

	var article models.Article
	var createdTime sql.NullTime

	// Scan メソッドを利用して、article と createdTime にデータを格納する
	err = row.Scan(&article.ID, &article.Title, &article.Contents,
		&article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		fmt.Println(err)
		return
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	fmt.Printf("%+v\n", article)
}
