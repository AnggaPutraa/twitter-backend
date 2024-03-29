// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: post.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (
  body,
  user_id
) VALUES (
  $1,
  $2
) RETURNING id, body, created_at, updated_at, user_id
`

type CreatePostParams struct {
	Body   string    `json:"body"`
	UserID uuid.UUID `json:"user_id"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost, arg.Body, arg.UserID)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Body,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
	)
	return i, err
}

const getAllPost = `-- name: GetAllPost :many
SELECT id, body, created_at, updated_at, user_id
FROM posts
`

func (q *Queries) GetAllPost(ctx context.Context) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getAllPost)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Body,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
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

const getPostById = `-- name: GetPostById :one
SELECT id, body, created_at, updated_at, user_id
FROM posts
WHERE id = $1
`

func (q *Queries) GetPostById(ctx context.Context, id uuid.UUID) (Post, error) {
	row := q.db.QueryRowContext(ctx, getPostById, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Body,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
	)
	return i, err
}

const updatePost = `-- name: UpdatePost :one
UPDATE posts
SET
  body = $2
WHERE
  id = $1 AND
  user_id = $3
RETURNING id, body, created_at, updated_at, user_id
`

type UpdatePostParams struct {
	ID     uuid.UUID `json:"id"`
	Body   string    `json:"body"`
	UserID uuid.UUID `json:"user_id"`
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, updatePost, arg.ID, arg.Body, arg.UserID)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Body,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
	)
	return i, err
}
