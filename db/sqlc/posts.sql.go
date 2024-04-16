// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: posts.sql

package db

import (
	"context"
	"time"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (
    content, author_id
) VALUES (
    $1, $2
) RETURNING id, content, author_id, created_at, updated_at
`

type CreatePostParams struct {
	Content  string `json:"content" validate:"required"`
	AuthorID int64  `json:"author_id" validate:"required"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost, arg.Content, arg.AuthorID)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.AuthorID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deletePost = `-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1
`

func (q *Queries) DeletePost(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePost, id)
	return err
}

const getPost = `-- name: GetPost :one
SELECT
    p.id,
    p.author_id,
    p.content,
    p.created_at,
    p.updated_at,
    lc.count_likes
FROM posts p
LEFT JOIN (
    SELECT post_id, COUNT(*) as count_likes
    FROM post_likes
    GROUP BY post_id
) lc ON p.id = lc.post_id
WHERE p.id = $1
LIMIT 1
`

type GetPostRow struct {
	ID         int64     `json:"id"`
	AuthorID   int64     `json:"author_id" validate:"required"`
	Content    string    `json:"content" validate:"required"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	CountLikes int64     `json:"count_likes"`
}

func (q *Queries) GetPost(ctx context.Context, id int64) (GetPostRow, error) {
	row := q.db.QueryRowContext(ctx, getPost, id)
	var i GetPostRow
	err := row.Scan(
		&i.ID,
		&i.AuthorID,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CountLikes,
	)
	return i, err
}

const getPostsAuthorID = `-- name: GetPostsAuthorID :one
SELECT p.author_id
FROM posts p
WHERE p.id = $1 LIMIT 1
`

func (q *Queries) GetPostsAuthorID(ctx context.Context, id int64) (int64, error) {
	row := q.db.QueryRowContext(ctx, getPostsAuthorID, id)
	var author_id int64
	err := row.Scan(&author_id)
	return author_id, err
}

const listUserPosts = `-- name: ListUserPosts :many
SELECT
    p.id,
    p.author_id,
    p.content,
    p.created_at,
    p.updated_at,
    lc.count_likes
FROM posts p
LEFT JOIN (
    SELECT post_id, COUNT(*) as count_likes
    FROM post_likes
    GROUP BY post_id
) lc ON p.id = lc.post_id
WHERE p.author_id = $1
ORDER BY p.id
`

type ListUserPostsRow struct {
	ID         int64     `json:"id"`
	AuthorID   int64     `json:"author_id" validate:"required"`
	Content    string    `json:"content" validate:"required"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	CountLikes int64     `json:"count_likes"`
}

func (q *Queries) ListUserPosts(ctx context.Context, authorID int64) ([]ListUserPostsRow, error) {
	rows, err := q.db.QueryContext(ctx, listUserPosts, authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListUserPostsRow{}
	for rows.Next() {
		var i ListUserPostsRow
		if err := rows.Scan(
			&i.ID,
			&i.AuthorID,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CountLikes,
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

const updatePost = `-- name: UpdatePost :one
UPDATE posts
SET content = $2
WHERE id = $1
RETURNING id, content, author_id, created_at, updated_at
`

type UpdatePostParams struct {
	ID      int64  `json:"id"`
	Content string `json:"content" validate:"required"`
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, updatePost, arg.ID, arg.Content)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.AuthorID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
