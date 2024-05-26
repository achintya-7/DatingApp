package app

import (
	"database/sql"
	"net/http"

	"github.com/achintya-7/dating-app/config"
	"github.com/achintya-7/dating-app/internal/controllers"
	"github.com/achintya-7/dating-app/internal/middleware"
	"github.com/achintya-7/dating-app/logger"
	"github.com/achintya-7/dating-app/pkg/mail"
	db "github.com/achintya-7/dating-app/pkg/sql/sqlc"
	"github.com/achintya-7/dating-app/pkg/token"
	distributor "github.com/achintya-7/dating-app/pkg/worker/distributor"
	processor "github.com/achintya-7/dating-app/pkg/worker/processor"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
)

type Server struct {
	router      *gin.Engine
	store       *db.Store
	tokenMaker  *token.PasetoMaker
	distributor distributor.TaskDistributor
	processor   *processor.RedisTaskProcessor
}

func NewServer() *Server {
	server := &Server{}

	server.setupDatbases()
	server.setupClient()
	server.setupWorker()
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

func (s *Server) setupWorker() {
	redisOpt := asynq.RedisClientOpt{
		Addr: config.Values.RedisUrl,
	}

	taskDistributor := distributor.NewRedisTaskDistributor(redisOpt)

	// Create a mail sender
	mailer := mail.NewGmailSender(config.Values.EmailName, config.Values.EmailAddress, config.Values.EmailPassowrd)

	// Create a processor instance
	redisTaskProcessor := processor.NewRedisTaskProcessor(redisOpt, s.store, mailer)

	s.distributor = taskDistributor
	s.processor = redisTaskProcessor
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
	v1Router := controllers.NewV1Router(s.store, s.tokenMaker)
	v1Router.SetupRoutes(baseRouter)

	// Setup v2 routes
	v2Router := controllers.NewV2Router(s.store, s.tokenMaker, s.distributor)
	v2Router.SetupRoutes(baseRouter)

	// Setup CORS middleware
	router.Use(middleware.CorsMiddleware())

	s.router = router
}

func (s *Server) Start() error {
	go func() {
		logger.Info(nil, "starting task processor")
		err := s.processor.Start()
		if err != nil {
			logger.Fatal(nil, "cannot start task processor")
		}
	}()

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
