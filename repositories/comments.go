package repositories

import (
	"database/sql"
	"go-myapi/models"
)

// 新規投稿をデータベースにinsertする関数
// ->データベースに保存したコメント内容と、発生したエラーを戻り値にする
func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `insert into comments (article_id, message, created_at) values (?, ?, now());`

	var newComment models.Comment
	newComment.ArticleID = comment.ArticleID
	newComment.Message = comment.Message

	result, err := db.Exec(sqlStr, comment.ArticleID, comment.Message)
	if err != nil {
		return models.Comment{}, err
	}
	id, _ := result.LastInsertId()
	newComment.CommentID = int(id)

	return newComment, nil
}

// 指定IDの記事についたコメント一覧を取得する関数
// -> 取得したコメントデータと、発生したエラーを戻り値にする
func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `select * from comments where article_id = ?;`

	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	commentArray := make([]models.Comment, 0)
	for rows.Next() {
		var comment models.Comment
		var createdTime sql.NullTime
		rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &createdTime)

		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}

		commentArray = append(commentArray, comment)
	}
	return commentArray, nil
}

// コメントを削除する関数
func DeleteComment(db *sql.DB, commentID int) error {
	const sqlStr = `
		delete from comments
		where comment_id = ?
	`
	_, err := db.Exec(sqlStr, commentID)
	if err != nil {
		return err
	}

	return nil
}
