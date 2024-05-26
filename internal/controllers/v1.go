package controllers

import (
	"github.com/achintya-7/dating-app/internal/dto"
	v1 "github.com/achintya-7/dating-app/internal/handlers/v1"
	db "github.com/achintya-7/dating-app/pkg/sql/sqlc"
	"github.com/achintya-7/dating-app/pkg/token"
	"github.com/achintya-7/dating-app/utils"
	"github.com/gin-gonic/gin"
)

type Router struct {
	handlers *v1.RouteHandler
}

func NewRouter(store *db.Store, tokenMaker *token.PasetoMaker) *Router {
	return &Router{
		handlers: v1.NewRouteHandler(store, tokenMaker),
	}
}

func (s *Router) SetupRoutes(router *gin.RouterGroup) {
	v1Route := router.Group("/v1")

	v1Route.POST("/login", utils.HandlerWrapper[dto.LoginResponse](s.handlers.Login))

	// Setup user routes
	usersRoute := v1Route.Group("/users")
	usersRoute.POST("/create", utils.HandlerWrapper[dto.CreateUserResponse](s.handlers.CreateUser))
}
