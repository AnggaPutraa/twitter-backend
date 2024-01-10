package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreatePostRequest struct {
	Body string `json:"body" binding:"required"`
}

type UpdatePostParams struct {
	Id string `uri:"id" binding:"required"`
}

type UpdatePostRequest struct {
	Body string `json:"body" binding:"required"`
}

type GetPostByIdParams struct {
	Id string `uri:"id" binding:"required"`
}

type PostResponse struct {
	ID        uuid.UUID         `json:"id"`
	Body      string            `json:"body"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	UserID    uuid.UUID         `json:"user_id"`
	Comments  []CommentResponse `json:"comments"`
}
