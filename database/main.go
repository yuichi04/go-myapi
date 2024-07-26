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

	// // sql.DB 型の Ping メソッドで疎通確認をする
	// if err := db.Ping(); err != nil {
	// 	// 失敗したらエラーを出力
	// 	fmt.Println(err)
	// } else {
	// 	// 成功したらメッセージを出力
	// 	fmt.Println("connect to DB")
	// }

	const sqlStr = `
		select * from articles;
	`

	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close() // Rows型もHTTPリクエストボディ同様「使い終わったらCloseしなければならない」類の構造体

	articleArray := make([]models.Article, 0)

	// rows に存在するレコードそれぞれに対して、繰り返し処理を実行する
	for rows.Next() {
		// 変数 article の各フィールドに、取得レコードのデータを入れる
		// （SQL クエリの select 句から、タイトル・本文・ユーザー名・いいね数が返ってくることはわかっている）
		var article models.Article
		var createdTime sql.NullTime
		err := rows.Scan(&article.ID, &article.Title, &article.Contents,
			&article.UserName, &article.NiceNum, &createdTime)

		if createdTime.Valid {
			article.CreatedAt = createdTime.Time
		}

		if err != nil {
			fmt.Println(err)
		} else {
			articleArray = append(articleArray, article)
		}
	}

	fmt.Printf("%+v\n", articleArray)
}
