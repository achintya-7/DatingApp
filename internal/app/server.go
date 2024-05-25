package app

import (
	"database/sql"
	"log"

	"github.com/achintya-7/dating-app/config"
	v1 "github.com/achintya-7/dating-app/internal/controllers"
	db "github.com/achintya-7/dating-app/pkg/sql/sqlc"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	store  *db.Store
}

func NewServer() *Server {
	server := &Server{}

	server.setupDatbases()
	server.setupRouter()

	return server
}

func (s *Server) setupDatbases() {
	conn, err := sql.Open("mysql", config.Values.MySqlUrl)
	if err != nil {
		log.Fatal(err)
	}

	store := db.NewStore(conn)

	s.store = store
}

func (s *Server) setupRouter() {
	router := gin.Default()

	baseRouter := router.Group("/dating-app")

	// Setup v1 routes
	v1Router := v1.NewRouter(s.store)
	v1Router.SetupRoutes(baseRouter)

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"*"},
	}))

	s.router = router
}

func (s *Server) Start() error {
	port := config.Values.HttpPort
	return s.router.Run(port)
}
