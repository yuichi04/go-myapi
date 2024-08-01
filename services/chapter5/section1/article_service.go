package services

import (
	"go-myapi/models"
	"go-myapi/repositories"
)

// ArticleDetailHandler用のサービス関数
func GetArticleService(articleID int) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	// 1. repositories層の関数SelectArticleDetailで記事の詳細を取得
	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	// 2. repositoeies層の関数SelectCommentListでコメント一覧を取得
	commentList, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	// 3. 2で得たコメント一覧を、1で得たArticle構造体に紐付ける
	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

// PostArticleHandler用のサービス関数
// 引数の情報を元に新しい記事を作り、結果を返却
func PostArticleService(article models.Article) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	newArticle, err := repositories.InsertArticle(db, article)
	if err != nil {
		return models.Article{}, err
	}
	return newArticle, nil
}

// ArticleListHandler用のサービス関数
// 指定のpageの記事一覧を返却
func GetArticleListService(page int) ([]models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return []models.Article{}, err
	}
	defer db.Close()

	articleList, err := repositories.SelectArticleList(db, page)
	if err != nil {
		return []models.Article{}, err
	}

	return articleList, nil
}

// PostNiceHandler用のサービス関数
// 指定IDの記事のいいね数を+1して、結果を返却
func PostNiceService(article models.Article) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	err = repositories.UpdateNiceNum(db, article.ID)
	if err != nil {
		return models.Article{}, err
	}

	updatedArticle, err := repositories.SelectArticleDetail(db, article.ID)
	if err != nil {
		return models.Article{}, err
	}

	return updatedArticle, nil
}

// SelectCommentList用のサービス関数
// 指定IDの記事のコメントリストを取得して、返す
func GetCommentListService(article models.Article) ([]models.Comment, error) {
	db, err := connectDB()
	if err != nil {
		return []models.Comment{}, err
	}

	commentList, err := repositories.SelectCommentList(db, article.ID)
	if err != nil {
		return []models.Comment{}, err
	}

	return commentList, nil
}
