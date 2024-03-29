// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: comment.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createComment = `-- name: CreateComment :one
INSERT INTO comments (
    body, 
    post_id, 
    user_id
) VALUES (
    $1, 
    $2, 
    $3
) RETURNING id, body, created_at, updated_at, user_id, post_id, parent_comment_id
`

type CreateCommentParams struct {
	Body   string    `json:"body"`
	PostID uuid.UUID `json:"post_id"`
	UserID uuid.UUID `json:"user_id"`
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error) {
	row := q.db.QueryRowContext(ctx, createComment, arg.Body, arg.PostID, arg.UserID)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.Body,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.PostID,
		&i.ParentCommentID,
	)
	return i, err
}

const createReplyComment = `-- name: CreateReplyComment :one
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
)RETURNING id, body, created_at, updated_at, user_id, post_id, parent_comment_id
`

type CreateReplyCommentParams struct {
	Body            string        `json:"body"`
	PostID          uuid.UUID     `json:"post_id"`
	UserID          uuid.UUID     `json:"user_id"`
	ParentCommentID uuid.NullUUID `json:"parent_comment_id"`
}

func (q *Queries) CreateReplyComment(ctx context.Context, arg CreateReplyCommentParams) (Comment, error) {
	row := q.db.QueryRowContext(ctx, createReplyComment,
		arg.Body,
		arg.PostID,
		arg.UserID,
		arg.ParentCommentID,
	)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.Body,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.PostID,
		&i.ParentCommentID,
	)
	return i, err
}

const getCommentByPost = `-- name: GetCommentByPost :many
SELECT id, body, created_at, updated_at, user_id, post_id, parent_comment_id
FROM comments
WHERE post_id = $1
`

func (q *Queries) GetCommentByPost(ctx context.Context, postID uuid.UUID) ([]Comment, error) {
	rows, err := q.db.QueryContext(ctx, getCommentByPost, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Comment{}
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.Body,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.PostID,
			&i.ParentCommentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getParentCommentByPost = `-- name: GetParentCommentByPost :many
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
  parent_comment_id IS NULL
`

func (q *Queries) GetParentCommentByPost(ctx context.Context, postID uuid.UUID) ([]Comment, error) {
	rows, err := q.db.QueryContext(ctx, getParentCommentByPost, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Comment{}
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.Body,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.PostID,
			&i.ParentCommentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateComment = `-- name: UpdateComment :one
UPDATE comments
SET
  body = $2
WHERE
  id = $1 AND
  user_id = $3
RETURNING id, body, created_at, updated_at, user_id, post_id, parent_comment_id
`

type UpdateCommentParams struct {
	ID     uuid.UUID `json:"id"`
	Body   string    `json:"body"`
	UserID uuid.UUID `json:"user_id"`
}

func (q *Queries) UpdateComment(ctx context.Context, arg UpdateCommentParams) (Comment, error) {
	row := q.db.QueryRowContext(ctx, updateComment, arg.ID, arg.Body, arg.UserID)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.Body,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.PostID,
		&i.ParentCommentID,
	)
	return i, err
}
