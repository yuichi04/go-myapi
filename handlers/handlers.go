package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
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

/*
  標準パッケージを使ったハンドラ関数
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
	// POSTメソッド以外は許可しない
	if req.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
	// POSTメソッドのリクエストに対してレスポンスを返す
	io.WriteString(w, "Posting Article...\n")
}

// GET /article/list のハンドラ
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	// GETメソッド以外は許可しない
	if req.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
	// GETメソッドのリクエストに対してレスポンスを返す
	io.WriteString(w, "Article List\n")
}

// GET /article/1 のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	// GETメソッド以外は許可しない
	if req.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
	// GETメソッドのリクエストに対してレスポンスを返す
	articleID := 1
	resStr := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resStr)
}

// POST /article/nice のハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	// POSTメソッド以外は許可しない
	if req.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
	// POSTメソッドのリクエストに対してレスポンスを返す
	io.WriteString(w, "Posting Nice...\n")
}

// POST /comment のハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	// POSTメソッド以外は許可しない
	if req.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
	// POSTメソッドのリクエストに対してレスポンスを返す
	io.WriteString(w, "Posting Comment...\n")
}

/*
Echoを使ったハンドラ関数
*/
func EchoHelloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World\n")
}

func EchoPostArticleHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Posting Article...\n")
}

func EchoArticleListHandler(c echo.Context) error {
	pageStr := c.QueryParam("page")
	if pageStr == "" {
		pageStr = "1" // デフォルトページを1に設定
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		return c.String(http.StatusBadRequest, "Invalid or missing 'page' query parameter\n")
	}
	resStr := fmt.Sprintf("Article List (page %d)\n", page)
	return c.String(http.StatusOK, resStr)
}

func EchoArticleDetailHandler(c echo.Context) error {
	articleId, err := strconv.Atoi(c.Param("articleId"))
	if err != nil || articleId < 0 {
		// errをそのまま返すと500番（Internal Server Error）になるため、400番（Bad Request）を返す
		// 負の数も不正なパラメータとして扱う
		return c.String(http.StatusBadRequest, "Invalid query parameter\n")
	}
	resStr := fmt.Sprintf("Article No.%v\n", articleId)

	return c.String(http.StatusOK, resStr)
}

func EchoPostNiceHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Posting Nice...\n")
}

func EchoPostCommentHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Posting Comment...\n")
}
