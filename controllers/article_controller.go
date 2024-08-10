package controllers

import (
	"encoding/json"
	"go-myapi/apperrors"
	"go-myapi/controllers/services"
	"go-myapi/models"
	"net/http"
	"strconv"
	"strings"
)

// Article用のコントローラ構造体
type ArticleController struct {
	service services.ArticleServicer
}

// コンストラクタ関数
func NewArticleController(s services.ArticleServicer) *ArticleController {
	return &ArticleController{service: s}
}

func (c *ArticleController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
	}

	article, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func (c *ArticleController) ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()
	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			err = apperrors.BadParam.Wrap(err, "queryparam must be number")
			apperrors.ErrorHandler(w, req, err)
			return
		}
	} else {
		page = 1
	}

	articleList, err := c.service.GetArticleListService(page)
	if err != nil {
		http.Error(w, "invalid query parameter", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(articleList)
}

func (c *ArticleController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	parts := strings.Split(req.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "invalid query parameter", http.StatusBadRequest)
		return
	}

	articleID, err := strconv.Atoi(parts[2])
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "queryparam must be number")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	article, err := c.service.GetArticleService(articleID)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func (c *ArticleController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	updatedArticle, err := c.service.PostNiceService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedArticle)
}
