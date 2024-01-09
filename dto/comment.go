package dto

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
