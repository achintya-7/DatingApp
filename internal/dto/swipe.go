package dto

import db "github.com/achintya-7/dating-app/pkg/sql/sqlc"

type SwipeRequest struct {
	SwipedUserId string             `json:"swiper_user_id" binding:"required"`
	Preference   db.SwipesSwipeType `json:"preference" binding:"required" enum:"YES,NO"`
}
