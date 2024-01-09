// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: user.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
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
) RETURNING id, name, username, email, hashed_password, bio, hashed_refresh_token, created_at, updated_at
`

type CreateUserParams struct {
	Name               string         `json:"name"`
	Username           string         `json:"username"`
	Email              string         `json:"email"`
	HashedPassword     string         `json:"hashed_password"`
	Bio                sql.NullString `json:"bio"`
	HashedRefreshToken sql.NullString `json:"hashed_refresh_token"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Name,
		arg.Username,
		arg.Email,
		arg.HashedPassword,
		arg.Bio,
		arg.HashedRefreshToken,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		&i.Bio,
		&i.HashedRefreshToken,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, name, username, email, hashed_password, bio, hashed_refresh_token, created_at, updated_at
FROM users
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		&i.Bio,
		&i.HashedRefreshToken,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, name, username, email, hashed_password, bio, hashed_refresh_token, created_at, updated_at
FROM users
WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		&i.Bio,
		&i.HashedRefreshToken,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateHashedRefreshToken = `-- name: UpdateHashedRefreshToken :one
UPDATE users
SET hashed_refresh_token =  $1
WHERE id = $2
RETURNING id, name, username, email, hashed_password, bio, hashed_refresh_token, created_at, updated_at
`

type UpdateHashedRefreshTokenParams struct {
	HashedRefreshToken sql.NullString `json:"hashed_refresh_token"`
	ID                 uuid.UUID      `json:"id"`
}

func (q *Queries) UpdateHashedRefreshToken(ctx context.Context, arg UpdateHashedRefreshTokenParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateHashedRefreshToken, arg.HashedRefreshToken, arg.ID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		&i.Bio,
		&i.HashedRefreshToken,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
