-- name: CreateUser :execresult
INSERT INTO Users (user_id, email, password, name, gender, age, latitude, longitude, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW()); 

-- name: GetUserByEmail :one
SELECT user_id, email, password, name
FROM Users
WHERE email = ?;

-- name: DiscoverUsersV1 :many
SELECT user_id, name, gender, age
FROM Users
WHERE user_id NOT IN (
    SELECT swipee_id
    FROM Swipes
    WHERE swiper_id = ?
);

-- name: DiscoverUsersV2 :many
SELECT user_id, name, gender, age
FROM Users
WHERE user_id NOT IN (
    SELECT swipee_id
    FROM Swipes
    WHERE swiper_id = $1
) 
AND age >= COALESCE(sqlc.narg(greater_than_age), age)
AND age <= COALESCE(sqlc.narg(lower_than_age), age)
AND gender = COALESCE(sqlc.narg(gender), gender);