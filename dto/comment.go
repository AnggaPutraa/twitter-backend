package dto

import "github.com/google/uuid"

type CreateCommentRequest struct {
	Body   string `json:"body" binding:"required"`
	PostID string `json:"post_id" binding:"required"`
}

type UpdateCommentParam struct {
	Id string `uri:"id" binding:"required"`
}

type UpdateCommentRequest struct {
	Body string `json:"body" binding:"required"`
}

type CreateReplyCommentRequest struct {
	Body            string `json:"body" binding:"required"`
	PostID          string `json:"post_id" binding:"required"`
	ParentCommentID string `json:"parent_comment_id" binding:"required"`
}

type CommentResponse struct {
	ID      uuid.UUID         `json:"id"`
	Body    string            `json:"body"`
	UserID  uuid.UUID         `json:"user_id"`
	PostID  uuid.UUID         `json:"post_id"`
	Replies []CommentResponse `json:"replies"`
}
