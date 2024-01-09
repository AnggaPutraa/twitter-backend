
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