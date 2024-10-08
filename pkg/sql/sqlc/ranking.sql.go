// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: ranking.sql

package db

import (
	"context"
	"database/sql"
)

const createRanking = `-- name: CreateRanking :execresult
INSERT INTO Rankings (user_id, like_count, dislike_count, attractiveness_score)
VALUES (?, 0, 0, 0.0)
`

func (q *Queries) CreateRanking(ctx context.Context, userID string) (sql.Result, error) {
	return q.db.ExecContext(ctx, createRanking, userID)
}

const getRankingByUserId = `-- name: GetRankingByUserId :one
SELECT user_id, like_count, dislike_count, attractiveness_score FROM Rankings WHERE user_id = ?
`

func (q *Queries) GetRankingByUserId(ctx context.Context, userID string) (Ranking, error) {
	row := q.db.QueryRowContext(ctx, getRankingByUserId, userID)
	var i Ranking
	err := row.Scan(
		&i.UserID,
		&i.LikeCount,
		&i.DislikeCount,
		&i.AttractivenessScore,
	)
	return i, err
}

const updateRanking = `-- name: UpdateRanking :execresult
UPDATE Rankings SET like_count = ?, dislike_count = ?, attractiveness_score = ? WHERE user_id = ?
`

type UpdateRankingParams struct {
	LikeCount           int32   `json:"like_count"`
	DislikeCount        int32   `json:"dislike_count"`
	AttractivenessScore float64 `json:"attractiveness_score"`
	UserID              string  `json:"user_id"`
}

func (q *Queries) UpdateRanking(ctx context.Context, arg UpdateRankingParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateRanking,
		arg.LikeCount,
		arg.DislikeCount,
		arg.AttractivenessScore,
		arg.UserID,
	)
}
