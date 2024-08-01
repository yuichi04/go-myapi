package handlers

import (
	"encoding/json"
	"go-myapi/models"
	"go-myapi/services"
	"io"
	"net/http"
	"strconv"
	"strings"
)

/*
  # ハンドラ:
  HTTPリクエストを受け取って、それに対するHTTPレスポンスの内容をコネクションに書き込む関数のこと

  ## ハンドラ関数の書き方の流れ:
  1. 第2引数 req*http.Request の中身を使って、レスポンスの中身を作成する
  2. 作成したレスポンスの中身を第一引数 w http.ResponseWriter に書き込む

  ## ハンドラの第2引数がポインタ型である理由:
  値渡しではなく参照渡しにすることで大きな構造体をコピーする必要がなくなり、パフォーマンスが向上するため
  ※値渡しの場合、関数に引数を渡す際にその引数のコピーを作成して渡すため、その分パフォーマンスが落ちてしまう

  ## ハンドラの第1引数がポインタ型ではない理由:
  http.ResponseWriterはインターフェースであり、既に参照型であるためポインタを使用する必要がないため
  ※ポインタ型を渡すことは「参照渡し」に該当する
*/

// GET /hello のハンドラ
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	// GETメソッド以外は許可しない
	if req.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
	// GETメソッドのリクエストに対してレスポンスを返す
	io.WriteString(w, "Hello, World!\n")
}

// POST /article のハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	// 1. json.NewDecoder(req.Body): req.Body（HTTPリクエストボディ）から読み取るための新しいデコーダを作成
	// 2. .Decode(&article): 作成したデコーダを使用して、JSONデータを`reqArticle`の変数にデコード（変換）
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}
	newArticle, err := services.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	// 1. json.NewEncoder(w):（HTTPレスポンスライター）に書き込むための新しいエンコーダを作成
	// 2. .Encode(article): 作成したエンコーダを使用して、articleの変数をJSONデータにエンコード（変換）し、それを`w`に書き込む
	json.NewEncoder(w).Encode(newArticle)
}

// GET /article/list のハンドラ
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	articleList, err := services.GetArticleListService(page)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(articleList)
}

// GET /article/1 のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	// 1. パスから、取得した記事のIDを得る
	parts := strings.Split(req.URL.Path, "/") // URLが`/article/123`の形式であることを想定
	if len(parts) < 3 {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	articleID, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	// 2. 指定IDの記事をデータベースから取得する
	article, err := services.GetArticleService(articleID)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	// 3. 結果をレスポンスに書き込む
	json.NewEncoder(w).Encode(article)
}

// POST /article/nice のハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	// POSTリクエストのボディから記事データを取得
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	// 2. リクエストから取得した記事のいいね数を増加
	updatedArticle, err := services.PostNiceService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	// 3. 更新した記事データをJSON変換して返却
	json.NewEncoder(w).Encode(updatedArticle)
}

// POST /comment のハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	newComment, err := services.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(newComment)
}
