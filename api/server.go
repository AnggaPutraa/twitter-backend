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

	apiGroup.GET("/post", server.getAllPost)
	apiGroup.GET("/post/:id", server.getPostById)

	authenticatedRoute := apiGroup.Group("/", authMiddleware(server.strategy))
	authenticatedRoute.Use()
	{
		userRoute := authenticatedRoute.Group("/user")
		userRoute.POST("/following/:id", server.createNewFollowing)
		userRoute.DELETE("/following/:id", server.deleteFollowing)
		userRoute.POST("/following", server.deleteFollowing)
		userRoute.GET("/following/:id", server.getAllUserFollowing)
		userRoute.GET("/follower/:id", server.getAllUserFollowers)

		postRoute := authenticatedRoute.Group("/post")
		postRoute.POST("/", server.createPost)
		postRoute.PUT("/:id", server.updatePost)

		commentRoute := authenticatedRoute.Group("/comment")
		commentRoute.POST("/", server.createComment)
		commentRoute.PUT("/:id", server.updateComment)
		commentRoute.POST("/reply", server.createReplyComment)
	}

	server.router = router
}

func (server *Server) start(address string) error {
	return server.router.Run(address)
}

func RunServer(config configs.Config, query db.Querier) {
	server, _ := NewServer(config, query)
	server.start(config.ServerAddress)
}
