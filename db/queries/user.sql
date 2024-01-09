-- name: CreateUser :one
INSERT INTO users (
  name,
  username,
  email,
  hashed_password,
  bio,
  hashed_refresh_token
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
) RETURNING *;

-- name: GetUserById :one
SELECT *
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;

-- name: UpdateHashedRefreshToken :one
UPDATE users
SET hashed_refresh_token =  sqlc.arg(hashed_refresh_token)
WHERE id = sqlc.arg(id)
RETURNING *;
