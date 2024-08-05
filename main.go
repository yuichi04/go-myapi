package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"go-myapi/controllers"
	"go-myapi/routers"
	"go-myapi/services"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_NAME")
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func main() {
	// 1. サーバー全体で使用するsql.DB型を1つ生成する
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	// 2. sql.DB型をもとに、サーバー全体で使用するサービス構造体MyAppServiceを1つ生成する
	ser := services.NewMyAppService(db)

	// 3. MyAppService型をもとに、サーバー全体で使用するコントローラ構造体MyAppControllerを1つ生成する
	con := controllers.NewMyAppController(ser)

	// 4. コントローラ型MyAppControllerのハンドラメソッドとパストの関連付けを行う
	r := routers.NewRouter(con)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
