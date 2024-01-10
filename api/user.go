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

func (server *Server) createNewFollowing(ctx *gin.Context) {
	var params dto.CreateUserFollowingByIdParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.ErrorResponse(err))
		return
	}
	userId, err := uuid.Parse(params.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.ErrorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*utils.JWTClaims)
	param := db.CreateUserFollowingParams{
		FollowerID: authPayload.Sub,
		FollowesID: userId,
	}
	following, err := server.query.CreateUserFollowing(ctx, param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, following)
}

func (server *Server) deleteFollowing(ctx *gin.Context) {
	var request dto.DeleteUserFollowingByIdParams
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.ErrorResponse(err))
		return
	}
	userId, err := uuid.Parse(request.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.ErrorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*utils.JWTClaims)
	param := db.DeleteUserFollowingParams{
		FollowerID: authPayload.Sub,
		FollowesID: userId,
	}
	err = server.query.DeleteUserFollowing(ctx, param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, param)
}

func (server *Server) getAllUserFollowers(ctx *gin.Context) {
	var params dto.GetPostByIdParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.ErrorResponse(err))
		return
	}
	userId, err := uuid.Parse(params.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.ErrorResponse(err))
		return
	}
	users, err := server.query.GetUserFollowers(ctx, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exceptions.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (server *Server) getAllUserFollowing(ctx *gin.Context) {
	var params dto.GetPostByIdParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.ErrorResponse(err))
		return
	}
	userId, err := uuid.Parse(params.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.ErrorResponse(err))
		return
	}
	users, err := server.query.GetUserFollowing(ctx, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exceptions.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, users)
}
