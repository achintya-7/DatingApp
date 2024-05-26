package controllers

import (
	v2 "github.com/achintya-7/dating-app/internal/handlers/v2"
	"github.com/achintya-7/dating-app/internal/middleware"
	db "github.com/achintya-7/dating-app/pkg/sql/sqlc"
	"github.com/achintya-7/dating-app/pkg/token"
	distributor "github.com/achintya-7/dating-app/pkg/worker/distributor"
	"github.com/achintya-7/dating-app/utils"
	"github.com/gin-gonic/gin"
)

type V2Router struct {
	handlers   *v2.RouteHandler
	tokenMaker *token.PasetoMaker
}

func NewV2Router(store *db.Store, tokenMaker *token.PasetoMaker, distributor distributor.TaskDistributor) *V2Router {
	return &V2Router{
		handlers:   v2.NewRouteHandler(store, tokenMaker, distributor),
		tokenMaker: tokenMaker,
	}
}

func (r *V2Router) SetupRoutes(router *gin.RouterGroup) {
	v2Route := router.Group("/v2")

	// Setup user routes
	usersRoute := v2Route.Group("/users")
	usersRoute.POST("/create", utils.HandlerWrapper(r.handlers.CreateUserV2))

	// Apply auth middleware
	v2Route.Use(middleware.AuthMiddleware(*r.tokenMaker))

	// User Discovery route V2
	usersRoute.POST("/discover", utils.HandlerWrapper(r.handlers.DiscoverV2))
}
