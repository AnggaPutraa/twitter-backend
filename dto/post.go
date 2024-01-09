package dto

type CreatePostRequest struct {
	Body string `json:"body" binding:"required"`
}

type UpdatePostParams struct {
	Id string `uri:"id" binding:"required"`
}

type UpdatePostRequest struct {
	Body string `json:"body" binding:"required"`
}
