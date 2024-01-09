package api

import (
	"github.com/AnggaPutraa/twitter-backend/configs"
	db "github.com/AnggaPutraa/twitter-backend/db/sqlc"
	"github.com/AnggaPutraa/twitter-backend/utils"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config   configs.Config
	query    db.Querier
	strategy utils.Strategy
	router   *gin.Engine
}

func NewServer(config configs.Config, query db.Querier) (*Server, error) {
	server := &Server{
		config: config,
		query:  query,
		strategy: utils.NewJWTStrategy(
			config.AccessTokenSecret,
			config.RefreshTokenSecret,
		),
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.Use(responseInterceptor())

	apiGroup := router.Group("/api")

	apiGroup.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	authRoute := apiGroup.Group("/auth")
	authRoute.POST("/register", server.register)
	authRoute.POST("/login", server.login)

	server.router = router
}

func (server *Server) start(address string) error {
	return server.router.Run(address)
}

func RunServer(config configs.Config, query db.Querier) {
	server, _ := NewServer(config, query)
	server.start(config.ServerAddress)
}
