-- name: CreateRanking :execresult
INSERT INTO Rankings (user_id, like_count, dislike_count, attractiveness_score)
VALUES (?, 0, 0, 0.0)
