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