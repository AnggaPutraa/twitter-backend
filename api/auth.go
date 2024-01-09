package api

import (
	"database/sql"
	"net/http"

	db "github.com/AnggaPutraa/twitter-backend/db/sqlc"
	"github.com/AnggaPutraa/twitter-backend/dto"
	"github.com/AnggaPutraa/twitter-backend/exceptions"
	"github.com/AnggaPutraa/twitter-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (server *Server) register(ctx *gin.Context) {
	var request dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.ErrorResponse(err))
		return
	}

	hashedPassword := utils.Hash(request.Password)

	createUserParam := db.CreateUserParams{
		Email:          request.Email,
		Username:       request.Username,
		Name:           request.Name,
		HashedPassword: hashedPassword,
		Bio:            sql.NullString{String: request.Bio, Valid: request.Bio != ""},
	}

	user, err := server.query.CreateUser(ctx, createUserParam)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, err)
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, exceptions.ErrorResponse(err))
		return
	}

	accessToken, refreshToken, err := server.strategy.GenerateToken(user.ID, user.Email)
	var response = &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	ctx.JSON(http.StatusOK, response)
}

func (server *Server) login(ctx *gin.Context) {
	var request dto.LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.ErrorResponse(err))
		return
	}
	user, err := server.query.GetUserByEmail(ctx, request.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		ctx.JSON(http.StatusInternalServerError, exceptions.ErrorResponse(err))
		return
	}
	if err := utils.ComparePasswordWithHashed(request.Password, user.HashedPassword); err != nil {
		ctx.JSON(http.StatusUnauthorized, exceptions.ErrorResponse(err))
		return
	}

	accessToken, refreshToken, err := server.strategy.GenerateToken(user.ID, user.Email)
	var response = &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	ctx.JSON(http.StatusOK, response)
}
