package utils

import (
	db "github.com/AnggaPutraa/twitter-backend/db/sqlc"
	"github.com/AnggaPutraa/twitter-backend/dto"
	"github.com/google/uuid"
)

func GetRepliesOnComment(comments []db.Comment, parentCommentID uuid.UUID) []dto.CommentResponse {
	replies := make([]dto.CommentResponse, 0)
	for _, comment := range comments {
		if comment.ParentCommentID.Valid && comment.ParentCommentID.UUID == parentCommentID {
			commentResponse := dto.CommentResponse{
				ID:      comment.ID,
				Body:    comment.Body,
				UserID:  comment.UserID,
				PostID:  comment.PostID,
				Replies: GetRepliesOnComment(comments, comment.ID),
			}
			replies = append(replies, commentResponse)
		}
	}
	return replies
}
