-- name: CreateUser :one
INSERT INTO users (
  name,
  username,
  email,
  hashed_password,
  bio
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
) RETURNING *;

-- name: GetUserById :one
SELECT *
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;

-- name: GetUserFollowers :many
SELECT 
  u.id,
  u.username,
  u.email
FROM users u
JOIN followers f ON u.id = f.follower_id
WHERE f.followes_id = $1;

-- name: GetUserFollowing :many
SELECT
  u.id,
  u.username,
  u.email
FROM users u
JOIN followers f ON u.id = f.followes_id
WHERE f.follower_id = $1;

-- name: CreateUserFollowing :one
INSERT INTO followers (
  follower_id, 
  followes_id
) VALUES (
  $1, 
  $2
) RETURNING *;

-- name: DeleteUserFollowing :exec
DELETE FROM followers
WHERE follower_id = $1 AND followes_id = $2;