package app

import (
	"database/sql"
	"net/http"

	"github.com/achintya-7/dating-app/config"
	v1 "github.com/achintya-7/dating-app/internal/controllers"
	"github.com/achintya-7/dating-app/internal/middleware"
	"github.com/achintya-7/dating-app/logger"
	db "github.com/achintya-7/dating-app/pkg/sql/sqlc"
	"github.com/achintya-7/dating-app/pkg/token"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router     *gin.Engine
	store      *db.Store
	tokenMaker *token.PasetoMaker
}

func NewServer() *Server {
	server := &Server{}

	server.setupDatbases()
	server.setupClient()
	server.setupRouter()

	return server
}

func (s *Server) setupDatbases() {
	conn, err := sql.Open("mysql", config.Values.MySqlUrl)
	if err != nil {
		logger.Fatal(nil, "cannot connect to database")
	}

	store := db.NewStore(conn)

	s.store = store
}

func (s *Server) setupClient() {
	tokenMaker, err := token.NewPasetoMaker(config.Values.TokenSymmetricKey)
	if err != nil {
		logger.Fatal(nil, "cannot create token maker")
	}

	s.tokenMaker = tokenMaker
}

func (s *Server) setupRouter() {
	router := gin.Default()

	// Create a base router group
	baseRouter := router.Group("/dating-app")

	// Register health route
	s.registerHealthRoute(baseRouter)

	// Setup middleware
	baseRouter.Use(middleware.SetCorrelationIdMiddleware())

	// Setup v1 routes
	v1Router := v1.NewRouter(s.store, s.tokenMaker)
	v1Router.SetupRoutes(baseRouter)

	// Setup CORS middleware
	router.Use(middleware.CorsMiddleware())

	s.router = router
}

func (s *Server) Start() error {
	port := config.Values.HttpPort
	port = ":" + port

	logger.Info(nil, "starting server at port "+port)
	return s.router.Run(port)
}

func (s *Server) registerHealthRoute(router *gin.RouterGroup) {
	router.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status": "up",
		})
	})
}
