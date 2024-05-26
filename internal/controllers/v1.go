package controllers

import (
	v1 "github.com/achintya-7/dating-app/internal/handlers/v1"
	db "github.com/achintya-7/dating-app/pkg/sql/sqlc"
	"github.com/achintya-7/dating-app/pkg/token"
	"github.com/achintya-7/dating-app/utils"
	"github.com/gin-gonic/gin"
)

type Router struct {
	handlers   *v1.RouteHandler
	tokenMaker *token.PasetoMaker
}

func NewRouter(store *db.Store, tokenMaker *token.PasetoMaker) *Router {
	return &Router{
		handlers:   v1.NewRouteHandler(store, tokenMaker),
		tokenMaker: tokenMaker,
	}
}

func (r *Router) SetupRoutes(router *gin.RouterGroup) {
	v1Route := router.Group("/v1")

	v1Route.POST("/login", utils.HandlerWrapper(r.handlers.Login))

	// Apply auth middleware
	// v1Route.Use(middleware.AuthMiddleware(*r.tokenMaker))

	// Setup user routes
	usersRoute := v1Route.Group("/users")
	usersRoute.POST("/create", utils.HandlerWrapper(r.handlers.CreateUser))

	// Setup match routes
	matchRoute := v1Route.Group("/match")
	matchRoute.POST("/match", utils.HandlerWrapper[gin.H](r.handlers.Match))
}
