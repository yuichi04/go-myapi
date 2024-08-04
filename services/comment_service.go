package services

import (
	"go-myapi/models"
	"go-myapi/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return newComment, nil
}

func (s *MyAppService) GetCommentListService(article models.Article) ([]models.Comment, error) {
	commentList, err := repositories.SelectCommentList(s.db, article.ID)
	if err != nil {
		return []models.Comment{}, err
	}

	return commentList, nil
}
