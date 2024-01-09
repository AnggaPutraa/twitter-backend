package api

import (
	"net/http"

	db "github.com/AnggaPutraa/twitter-backend/db/sqlc"
	"github.com/AnggaPutraa/twitter-backend/dto"
	"github.com/AnggaPutraa/twitter-backend/exceptions"
	"github.com/AnggaPutraa/twitter-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (server *Server) createComment(ctx *gin.Context) {
	var request dto.CreateCommentRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.ErrorResponse(err))
		return
	}
	postId, err := uuid.Parse(request.PostID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.ErrorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*utils.JWTClaims)

	param := db.CreateCommentParams{
		Body:   request.Body,
		PostID: postId,
		UserID: authPayload.Sub,
	}
	comment, err := server.query.CreateComment(ctx, param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exceptions.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, comment)
}

func (server *Server) updateComment(ctx *gin.Context) {
	var params dto.UpdateCommentParam
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.ErrorResponse(err))
		return
	}
	var request dto.UpdateCommentRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.ErrorResponse(err))
		return
	}
	commentId, err := uuid.Parse(params.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.ErrorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*utils.JWTClaims)
	param := db.UpdateCommentParams{
		ID:     commentId,
		Body:   request.Body,
		UserID: authPayload.Sub,
	}
	comment, err := server.query.UpdateComment(ctx, param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exceptions.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, comment)
}

func (server *Server) createReplyComment(ctx *gin.Context) {
	var request dto.CreateReplyCommentRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.ErrorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*utils.JWTClaims)
	postId, err := uuid.Parse(request.PostID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.ErrorResponse(err))
		return
	}
	parentCommentId, err := uuid.Parse(request.ParentCommentID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.ErrorResponse(err))
		return
	}
	param := db.CreateReplyCommentParams{
		Body:            request.Body,
		PostID:          postId,
		ParentCommentID: uuid.NullUUID{UUID: parentCommentId, Valid: true},
		UserID:          authPayload.Sub,
	}
	comment, err := server.query.CreateReplyComment(ctx, param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exceptions.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, comment)
}
