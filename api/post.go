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

func (server *Server) createPost(ctx *gin.Context) {
	var request dto.CreatePostRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.ErrorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*utils.JWTClaims)
	param := db.CreatePostParams{
		Body:   request.Body,
		UserID: authPayload.Sub,
	}
	post, err := server.query.CreatePost(ctx, param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exceptions.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, post)
}

func (server *Server) updatePost(ctx *gin.Context) {
	var params dto.UpdatePostParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.ErrorResponse(err))
		return
	}
	postId, err := uuid.Parse(params.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.ErrorResponse(err))
		return
	}
	var request dto.UpdatePostRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.ErrorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*utils.JWTClaims)
	param := db.UpdatePostParams{
		ID:     postId,
		Body:   request.Body,
		UserID: authPayload.Sub,
	}
	post, err := server.query.UpdatePost(ctx, param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exceptions.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, post)
}
