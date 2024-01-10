-- name: CreatePost :one
INSERT INTO posts (
  body,
  user_id
) VALUES (
  $1,
  $2
) RETURNING *;

-- name: UpdatePost :one
UPDATE posts
SET
  body = $2
WHERE
  id = $1 AND
  user_id = $3
RETURNING *;

-- name: GetPostById :one
SELECT *
FROM posts
WHERE id = $1;