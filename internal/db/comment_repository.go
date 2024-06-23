package db

import (
	"context"
	_ "embed"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var (
	//go:embed cypher/create_comment.cypher
	createCommentCypher string

	//go:embed cypher/create_like_on_comment.cypher
	createLikeOnCommentCypher string

	//go:embed cypher/update_comment.cypher
	updateCommentCypher string

	//go:embed cypher/delete_comment.cypher
	deleteCommentCypher string

	//go:embed cypher/delete_comment_like.cypher
	deleteCommentLikeCypher string

	//go:embed cypher/list_post_comments.cypher
	listPostComments string
)

type ICommentRepository interface {
	CreateComment(ctx context.Context, params CreateCommentParams) (comment Comment, err error)
	CreateCommentLike(ctx context.Context, params CreateCommentLikeParams) (err error)
	ListPostComments(ctx context.Context, id int64) (posts []Comment, err error)
	UpdateComment(ctx context.Context, params UpdateCommentParams) (comment Comment, err error)
	DeleteComment(ctx context.Context, userId, commentId int64) (err error)
	DeleteCommentLike(ctx context.Context, params DeleteCommentLikeParams) error
}

type CommentRepository struct {
	Reposytory[Comment]
}

func NewCommentRepository(driver neo4j.DriverWithContext) *CommentRepository {
	return &CommentRepository{
		Reposytory: Reposytory[Comment]{
			driver: driver,
		},
	}
}

func (cr *CommentRepository) CreateComment(ctx context.Context, params CreateCommentParams) (comment Comment, err error) {
	return cr.Create(ctx, params, createCommentCypher)
}

func (cr *CommentRepository) CreateCommentLike(ctx context.Context, params CreateCommentLikeParams) (err error) {
	return cr.Write(ctx, createLikeOnCommentCypher, params)
}

func (cr *CommentRepository) ListPostComments(ctx context.Context, id int64) (posts []Comment, err error) {
	return cr.List(ctx, listPostComments, id)
}

func (cr *CommentRepository) UpdateComment(ctx context.Context, params UpdateCommentParams) (comment Comment, err error) {
	return cr.Update(ctx, updateCommentCypher, params)
}

func (cr *CommentRepository) DeleteComment(ctx context.Context, userId, commentId int64) (err error) {
	return cr.Delete(ctx, deleteCommentCypher, map[string]interface{}{
		"id":        commentId,
		"author_id": userId,
	})
}

func (cr *CommentRepository) DeleteCommentLike(ctx context.Context, params DeleteCommentLikeParams) error {
	return cr.Delete(ctx, deleteCommentLikeCypher, params)
}