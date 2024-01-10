
-- name: CreateComment :one
INSERT INTO comments (
    body, 
    post_id, 
    user_id
) VALUES (
    $1, 
    $2, 
    $3
) RETURNING *;

-- name: UpdateComment :one
UPDATE comments
SET
  body = $2
WHERE
  id = $1 AND
  user_id = $3
RETURNING *;

-- name: CreateReplyComment :one
INSERT INTO comments (
    body, 
    post_id, 
    user_id,
    parent_comment_id
) VALUES (
    $1, 
    $2, 
    $3, 
    $4
)RETURNING *;

-- name: GetParentCommentByPost :many
SELECT
  id,
  body,
  created_at,
  updated_at,
  user_id,
  post_id,
  parent_comment_id
FROM
  comments
WHERE
  post_id = $1 AND
  parent_comment_id IS NULL;

-- name: GetCommentByPost :many
SELECT *
FROM comments
WHERE post_id = $1;
