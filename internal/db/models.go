// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"time"
)

type Comment struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content" validate:"required"`
	AuthorID  int64     `json:"author_id" validate:"required"`
	PostID    int64     `json:"post_id" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommentInfo struct {
	ID        int64     `json:"id"`
	AuthorID  int64     `json:"author_id"`
	PostID    int64     `json:"post_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Likes     int64     `json:"likes"`
}

type CommentLike struct {
	UserID    int64 `json:"user_id" validate:"required"`
	CommentID int64 `json:"comment_id" validate:"required"`
}

type Post struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content" validate:"required"`
	AuthorID  int64     `json:"author_id" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostInfo struct {
	ID        int64     `json:"id"`
	AuthorID  int64     `json:"author_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Likes     int64     `json:"likes"`
}

type PostLike struct {
	UserID int64 `json:"user_id" validate:"required"`
	PostID int64 `json:"post_id" validate:"required"`
}

type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email" validate:"required,email"`
	FirstName string    `json:"first_name" validate:"required"`
	LastName  string    `json:"last_name" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}

type UsersSubscription struct {
	UserID         int64 `json:"user_id"`
	SubscriptionID int64 `json:"subscription_id"`
}
