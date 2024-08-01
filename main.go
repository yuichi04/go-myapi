package main

import (
	"go-myapi/handlers"
	"log"
	"net/http"
)

func main() {
	/*
		r := echo.New()
	*/

	/*
	  標準パッケージを使ったルーティング
	*/
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.ArticleListHandler)
	http.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	http.HandleFunc("/article/nice", handlers.PostNiceHandler)
	http.HandleFunc("/comment", handlers.PostCommentHandler)

	/*
	  Echoを使ったルーティング
	*/

	/*
		r.GET("/echo-hello", handlers.EchoHelloHandler)
		r.POST("/echo-article", handlers.EchoPostArticleHandler)
		r.GET("/echo-article/list", handlers.EchoArticleListHandler)
		r.GET("/echo-article/:articleId", handlers.EchoArticleDetailHandler)
		r.POST("/echo-article/nice", handlers.EchoPostNiceHandler)
		r.POST("/echo-comment", handlers.EchoPostCommentHandler)
	*/

	/*
	  サーバ起動
	*/
	// 標準パッケージ
	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	// Echo
	/*
		log.Println("echo server start at port 8080")
		log.Fatal(r.Start(":8080"))
	*/
}
