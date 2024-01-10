package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

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

func (server *Server) refresh(ctx *gin.Context) {
	authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
	if len(authorizationHeader) == 0 {
		err := errors.New("Authorization header is not provided")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}

	fields := strings.Fields(authorizationHeader)
	if len(fields) < 2 {
		err := errors.New("Authorization header is not provided")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}

	authorizationType := strings.ToLower(fields[0])
	if authorizationType != authorizationTypeBearer {
		err := fmt.Errorf("unsupported authorization type %s", authorizationType)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}
	refreshToken := fields[1]
	payload, err := server.strategy.VerifyToken(refreshToken, utils.RefreshTokenType)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}

	accessToken, refreshToken, err := server.strategy.GenerateToken(payload.Sub, payload.Email)
	var response = &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	ctx.JSON(http.StatusOK, response)
}
