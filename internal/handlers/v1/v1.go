package v1

import db "github.com/achintya-7/dating-app/pkg/sql/sqlc"

type RouteHandler struct {
	store *db.Store
}

func NewRouteHandler(store *db.Store) *RouteHandler {
	return &RouteHandler{
		store: store,
	}
}
