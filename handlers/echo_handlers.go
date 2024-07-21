package handlers

import (
	"go-myapi/helpers"
	"go-myapi/models"
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

func EchoHelloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World\n")
}

func EchoPostArticleHandler(c echo.Context) error {
	var article models.Article
	if err := c.Bind(&article); err != nil {
		return helpers.ReturnErrorInJSONPretty(c, http.StatusBadRequest, "Invalid request body")
	}

	if err := c.JSONPretty(http.StatusOK, article, "    "); err != nil {
		return helpers.ReturnErrorInJSONPretty(c, http.StatusInternalServerError, "Fail to encode json")
	}

	return nil
}

func EchoArticleListHandler(c echo.Context) error {
	pageStr := c.QueryParam("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		return helpers.ReturnErrorInJSONPretty(c, http.StatusBadRequest, "Invalid or missing 'page' query parameter")
	}
	articles := []models.Article{models.Article1, models.Article2}
	// JSON()とJSONPretty()の機能的な違いは整形するか否か
	// JSONPretty()はデバッグ目的で使用することが推奨（整形する分、オーバーヘッドが増えるため）
	if err := c.JSONPretty(http.StatusOK, articles, "    "); err != nil {
		return helpers.ReturnErrorInJSONPretty(c, http.StatusInternalServerError, "Fail to encode json")
	}
	return nil
}

func EchoArticleDetailHandler(c echo.Context) error {
	articleId, err := strconv.Atoi(c.Param("articleId"))
	if err != nil || articleId < 0 {
		return helpers.ReturnErrorInJSONPretty(c, http.StatusBadRequest, "Invalid or missing 'articleId' query parameter")
	}
	if err := c.JSONPretty(http.StatusOK, models.Article1, "    "); err != nil {
		return helpers.ReturnErrorInJSONPretty(c, http.StatusInternalServerError, "Fail to encode json")
	}
	return nil
}

func EchoPostNiceHandler(c echo.Context) error {
	if err := c.JSONPretty(http.StatusOK, models.Article1, "    "); err != nil {
		return helpers.ReturnErrorInJSONPretty(c, http.StatusInternalServerError, "Fail to encode json")
	}
	return nil
}

func EchoPostCommentHandler(c echo.Context) error {
	if err := c.JSONPretty(http.StatusOK, models.Comment1, "    "); err != nil {
		return helpers.ReturnErrorInJSONPretty(c, http.StatusInternalServerError, "Fail to encode json")
	}
	return nil
}
