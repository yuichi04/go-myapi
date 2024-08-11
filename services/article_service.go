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
	"sync"
)

func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	// 記事の詳細を格納するための変数を宣言
	var article models.Article
	// コメントリストを格納するための変数を宣言
	var commentList []models.Comment
	// 記事取得エラーとコメント取得エラーを格納するための変数を宣言
	var articleGetErr, commentGetErr error

	// 記事取得時の排他制御用のミューテックスを宣言
	var amu sync.Mutex
	// コメント取得時の排他制御用のミューテックスを宣言
	var cmu sync.Mutex

	// ゴールーチンの完了を待つためのWaitGroupを宣言
	var wg sync.WaitGroup

	// 2つのゴールーチンを待つことをWaitGroupに通知
	wg.Add(2)

	// 記事の詳細を取得するゴールーチンを開始
	go func(db *sql.DB, articleID int) {
		// ゴールーチンが終了したことをWaitGroupに通知
		defer wg.Done()
		// 記事取得時の排他制御を開始
		amu.Lock()
		// 記事の詳細を取得し、変数に格納
		article, articleGetErr = repositories.SelectArticleDetail(s.db, articleID)
		// 記事取得時の排他制御を終了
		amu.Unlock()
	}(s.db, articleID)
	// 記事取得エラーが発生した場合の処理
	if articleGetErr != nil {
		// エラーがデータなしの場合の処理
		if errors.Is(articleGetErr, sql.ErrNoRows) {
			err := apperrors.NAData.Wrap(articleGetErr, "no data")
			return models.Article{}, err
		}
		// その他のエラーの場合の処理
		err := apperrors.GetDataFailed.Wrap(articleGetErr, "fail to get data")
		return models.Article{}, err
	}

	// コメントリストを取得するゴールーチンを開始
	go func(db *sql.DB, articleID int) {
		// ゴールーチンが終了したことをWaitGroupに通知
		defer wg.Done()
		// コメント取得時の排他制御を開始
		cmu.Lock()
		// コメントリストを取得し、変数に格納
		commentList, commentGetErr = repositories.SelectCommentList(s.db, articleID)
		// コメント取得時の排他制御を終了
		cmu.Unlock()
	}(s.db, articleID)
	// コメント取得エラーが発生した場合の処理
	if commentGetErr != nil {
		err := apperrors.GetDataFailed.Wrap(commentGetErr, "fail to get data")
		return models.Article{}, err
	}

	// すべてのゴールーチンが完了するのを待つ
	wg.Wait()

	// 取得したコメントリストを記事に追加
	article.CommentList = append(article.CommentList, commentList...)

	// 記事の詳細を返す
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
