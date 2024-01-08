// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"database/sql"
	"time"
)

type Comment struct {
	ID              int64         `json:"id"`
	Body            string        `json:"body"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       sql.NullTime  `json:"updated_at"`
	UserID          int64         `json:"user_id"`
	PostID          int64         `json:"post_id"`
	ParentCommentID sql.NullInt64 `json:"parent_comment_id"`
}

type Follower struct {
	FollowerID int64 `json:"follower_id"`
	FollowesID int64 `json:"followes_id"`
}

type Notification struct {
	ID        int64     `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UserID    int64     `json:"user_id"`
}

type Post struct {
	ID        int64        `json:"id"`
	Body      string       `json:"body"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	UserID    int64        `json:"user_id"`
}

type PostLike struct {
	UserID int64 `json:"user_id"`
	PostID int64 `json:"post_id"`
}

type User struct {
	ID                 int64          `json:"id"`
	Name               string         `json:"name"`
	Username           string         `json:"username"`
	Email              string         `json:"email"`
	HashedPassword     string         `json:"hashed_password"`
	Bio                sql.NullString `json:"bio"`
	HashedRefreshToken sql.NullString `json:"hashed_refresh_token"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          sql.NullTime   `json:"updated_at"`
}