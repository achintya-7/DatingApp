package dto

import db "github.com/achintya-7/dating-app/pkg/sql/sqlc"

type MatchRequest struct {
	UserId     string             `json:"user_id" binding:"required"`
	Preference db.SwipesSwipeType `json:"preference" binding:"required" enum:"YES,NO"`
}

type MatchedResponse struct {
	Matched bool `json:"matched"`
	MatchId string `json:"match_id"`
}

type UnMatchedResponse struct {
	Matched bool `json:"matched"`
}