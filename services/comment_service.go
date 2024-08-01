package services

import (
	"go-myapi/models"
	"go-myapi/repositories"
)

// PostCommentHandler用のサービス関数
// 引数の情報を元に新しいコメントを作り、結果を返却
func PostCommentService(comment models.Comment) (models.Comment, error) {
	db, err := connectDB()
	if err != nil {
		return models.Comment{}, err
	}
	defer db.Close()

	newComment, err := repositories.InsertComment(db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return newComment, nil
}
