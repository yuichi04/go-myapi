/*
並行処理を使用する動機
→処理を並行に実行できた方がレスポンスを速くできるから

並行処理にする箇所の確認方法
- 同時に実行しても問題ない場所
- A->Bの順で実装されているが、実はBがAを待たなくて良い場所

ゴールーチン
ゴルーチンを用意するにはgo文が必要。
ただし、go文の後に記述する関数は戻り値なしのものでなくてはならない

チャネル
異なるゴールーチン同士で値の送受信をするためのパイプを表す変数の型

チャネルによる送受信の特徴
- チャネルに値を送信するゴールーチンは、受信側が値を受信する準備ができるまで待ちの状態になる
- チャネルから値を受信するゴールーチンは、送信側が値を送信する準備ができるまで待ちの状態になる
→「送信側による送信処理」と「受信側による受信処理」は基本的には同時に行われる。
*/
package services

import (
	"database/sql"
	"errors"
	"go-myapi/apperrors"
	"go-myapi/models"
	"go-myapi/repositories"
)

func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	// 1. 変数の初期化
	var article models.Article
	var commentList []models.Comment
	var articleGetErr, commentGetErr error

	// 2. チャネルの作成
	type articleResult struct {
		article models.Article
		err     error
	}
	// 記事用のチャネル
	articleChan := make(chan articleResult)
	defer close(articleChan)

	// 3. ゴールーチンの起動
	// ch には articleChan が渡ってくる
	go func(ch chan<- articleResult, db *sql.DB, articleID int) {
		article, err := repositories.SelectArticleDetail(s.db, articleID)
		ch <- articleResult{article: article, err: err}
	}(articleChan, s.db, articleID)

	// 2. チャネルの作成
	type commentResult struct {
		commentList *[]models.Comment
		err         error
	}
	// コメント用のチャネル
	commentChan := make(chan commentResult)
	defer close(commentChan)

	// 3. ゴールーチンの起動
	// ch には commentChan が渡ってくる
	go func(ch chan<- commentResult, db *sql.DB, articleID int) {
		commentList, err := repositories.SelectCommentList(s.db, articleID)
		ch <- commentResult{commentList: &commentList, err: err}
	}(commentChan, s.db, articleID)

	// 4. 結果の受け取り
	// いずれかのcaseが2回実行されたら次の処理に移る
	// ただし、実質、各caseは1回ずつしか実行されないようになっているため、すべてのcaseが実行されたら次の処理に移る
	for i := 0; i < 2; i++ {
		// select文：複数のチャネル操作を同時に待つための構文
		select {
		case ar := <-articleChan:
			article, articleGetErr = ar.article, ar.err
		case cr := <-commentChan:
			commentList, commentGetErr = *cr.commentList, cr.err
		}
	}

	// 5. エラーチェック
	if articleGetErr != nil {
		if errors.Is(articleGetErr, sql.ErrNoRows) {
			err := apperrors.NAData.Wrap(articleGetErr, "no data")
			return models.Article{}, err
		}
		err := apperrors.GetDataFailed.Wrap(articleGetErr, "fail to get data")
		return models.Article{}, err
	}

	if commentGetErr != nil {
		err := apperrors.GetDataFailed.Wrap(commentGetErr, "fail to get data")
		return models.Article{}, err
	}

	// 6. 結果の統合
	article.CommentList = append(article.CommentList, commentList...)

	// 7. 結果の返却
	return article, nil
}

func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Article{}, err
	}

	return newArticle, nil
}

func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}

	if len(articleList) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}

	return articleList, nil
}

func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {
	err := repositories.UpdateNiceNum(s.db, article.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "does not exist target article")
			return models.Article{}, err
		}

		err = apperrors.UpdateDataFailed.Wrap(err, "fail to update nice count")
		return models.Article{}, err
	}

	return models.Article{
		ID:        article.ID,
		Title:     article.Title,
		Contents:  article.Contents,
		UserName:  article.UserName,
		NiceNum:   article.NiceNum + 1,
		CreatedAt: article.CreatedAt,
	}, nil
}
