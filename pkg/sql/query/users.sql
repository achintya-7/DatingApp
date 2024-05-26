-- name: CreateUser :execresult
INSERT INTO Users (user_id, email, password, name, gender, age, latitude, longitude, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW()); 

-- name: GetUserByEmail :one
SELECT user_id, email, password, name
FROM Users
WHERE email = ?;