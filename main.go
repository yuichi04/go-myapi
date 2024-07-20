package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	/*
	  ハンドラの宣言
	  ハンドラとは「HTTPリクエストを受け取って、それに対するHTTPレスポンスの内容をコネクションに書き込む関数」のこと
	  helloHandler変数の型は、関数型 func(w http.ResponseWriter, req *http.Request)
	*/
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		/*
		  ハンドラ関数の書き方の流れ
		  1. 第2引数 req *http.Request の中身を使って、レスポンスの中身を作成する　※この処理では単にHello, world!を返しているだけなのでスキップ
		  2. 作成したレスポンスの中身を、第一引数 w http.ResponseWriter に書き込む
		*/
		io.WriteString(w, "Hello, world!\n")
	}

	postArticleHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Article...\n")
	}

	postNiceHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Nice...\n")
	}

	postCommentHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Comment...\n")
	}

	articleListHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Article List\n")
	}

	articleDetailHandler := func(w http.ResponseWriter, req *http.Request) {
		articleID := 1
		resString := fmt.Sprintf("Article No.%d\n", articleID)
		io.WriteString(w, resString)
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/article", postArticleHandler)
	http.HandleFunc("/article/list", articleListHandler)
	http.HandleFunc("/article/1", articleDetailHandler)
	http.HandleFunc("/article/nice", postNiceHandler)
	http.HandleFunc("/comment", postCommentHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
