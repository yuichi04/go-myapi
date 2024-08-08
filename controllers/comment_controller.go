package controllers

import (
	"encoding/json"
	"go-myapi/controllers/services"
	"go-myapi/models"
	"net/http"
)

type CommentController struct {
	service services.CommentServicer
}

func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Response) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(reqComment); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	newComment, err := c.service.PostcommentService(reqComment)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newComment)
}
