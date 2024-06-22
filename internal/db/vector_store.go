package db

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/malyshEvhen/meow_mingle/internal/errors"
	"github.com/malyshEvhen/meow_mingle/internal/utils"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/db"
)

var (
	//go:embed cypher/create_user.cypher
	createUserCypher string

	//go:embed cypher/match_user_by_id.cypher
	getUserCypher string

	//go:embed cypher/create_post.cypher
	createPostCypher string

	//go:embed cypher/match_post_by_id.cypher
	getPostCypher string

	//go:embed cypher/create_comment.cypher
	createCommentCypher string
)

type VStore struct {
	driver neo4j.DriverWithContext
}

func NewVstore(d neo4j.DriverWithContext) IStore {
	return &VStore{
		driver: d,
	}
}

func (s *VStore) CreateUserTx(ctx context.Context, userForm CreateUserParams) (user User, execErr error) {
	session := s.driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	if _, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := persist[User](ctx, tx, createUserCypher, userForm)
		if err != nil {
			return nil, err
		}
		user = result
		return result, nil
	}); err != nil {
		execErr = err
		return
	}
	return
}

func (s *VStore) GetUserTx(ctx context.Context, id int64) (user GetUserRow, execErr error) {
	session := s.driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	params := map[string]interface{}{
		"id": id,
	}

	if _, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := retrieve[GetUserRow](ctx, tx, getUserCypher, params)
		if err != nil {
			return nil, err
		}
		user = result

		return result, nil
	}); err != nil {
		execErr = err
		return
	}

	return
}

func (s *VStore) CreatePostTx(ctx context.Context, params CreatePostParams) (post Post, execErr error) {
	session := s.driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	if _, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := persist[Post](ctx, tx, createPostCypher, params)
		if err != nil {
			return nil, err
		}
		post = result

		return result, nil
	}); err != nil {
		execErr = err
		return
	}
	return
}

func (s *VStore) GetPostTx(ctx context.Context, id int64) (post PostInfo, execErr error) {
	session := s.driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	params := map[string]interface{}{
		"id": id,
	}

	if _, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := retrieve[PostInfo](ctx, tx, getPostCypher, params)
		if err != nil {
			return nil, err
		}
		post = result

		return result, nil
	}); err != nil {
		execErr = err
		return
	}
	return
}

func (s *VStore) CreateCommentTx(ctx context.Context, params CreateCommentParams) (comment Comment, execErr error) {
	session := s.driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	if _, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := persist[Comment](ctx, tx, createCommentCypher, params)
		if err != nil {
			return nil, err
		}
		comment = result

		return result, nil
	}); err != nil {
		execErr = err
		return
	}
	return
}

func (s *VStore) CreatePostLikeTx(ctx context.Context, params CreatePostLikeParams) error {
	panic("not implemented") // TODO: Implement
}

func (s *VStore) CreateCommentLikeTx(ctx context.Context, params CreateCommentLikeParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *VStore) CreateSubscriptionTx(ctx context.Context, params CreateSubscriptionParams) error {
	panic("not implemented") // TODO: Implement
}

func (s *VStore) GetFeed(ctx context.Context, userId int64) (feed []PostInfo, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *VStore) ListUserPostsTx(ctx context.Context, userId int64) (posts []PostInfo, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *VStore) ListPostCommentsTx(ctx context.Context, id int64) (posts []CommentInfo, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *VStore) UpdatePostTx(ctx context.Context, userId int64, params UpdatePostParams) (post Post, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *VStore) UpdateCommentTx(ctx context.Context, userId int64, params UpdateCommentParams) (comment Comment, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *VStore) DeletePostTx(ctx context.Context, userId int64, postId int64) error {
	panic("not implemented") // TODO: Implement
}

func (s *VStore) DeletePostLikeTx(ctx context.Context, params DeletePostLikeParams) error {
	panic("not implemented") // TODO: Implement
}

func (s *VStore) DeleteCommentTx(ctx context.Context, userId int64, commentId int64) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *VStore) DeleteCommentLikeTx(ctx context.Context, params DeleteCommentLikeParams) error {
	panic("not implemented") // TODO: Implement
}

func (s *VStore) DeleteSubscriptionTx(ctx context.Context, params DeleteSubscriptionParams) error {
	panic("not implemented") // TODO: Implement
}

func persist[T any](
	ctx context.Context,
	tx neo4j.ManagedTransaction,
	cypher string,
	form any,
) (obj T, err error) {
	fail := func(msg string) (res T, err error) {
		err = errors.NewDatabaseError(fmt.Errorf("error occurred while %s", msg))
		return
	}

	params, err := mapToProperties(form)
	if err != nil {
		return fail(fmt.Sprintf("convert to properties: %s", err.Error()))
	}

	result, err := tx.Run(ctx, cypher, params)
	if err != nil {
		return fail(fmt.Sprintf("executing transaction: %s", err.Error()))
	}

	record, err := result.Single(ctx)
	if err != nil {
		return fail(fmt.Sprintf("saving new record: %s", err.Error()))
	}

	return parceRecord[T](record)
}

func retrieve[T any](
	ctx context.Context,
	tx neo4j.ManagedTransaction,
	cypher string,
	params map[string]interface{},
) (obj T, err error) {
	fail := func(msg string) (res T, err error) {
		err = errors.NewDatabaseError(fmt.Errorf("error occurred while %s", msg))
		return
	}

	result, err := tx.Run(ctx, cypher, params)
	if err != nil {
		return fail(fmt.Sprintf("executing transaction: %s", err.Error()))
	}

	record, err := result.Single(ctx)
	if err != nil {
		return fail(fmt.Sprintf("retrieving the record: %s", err.Error()))
	}

	return parceRecord[T](record)
}

func parceRecord[T any](r *db.Record) (result T, err error) {
	fail := func(msg string) (res T, err error) {
		err = errors.NewDatabaseError(fmt.Errorf("error occurred while %s", msg))
		return
	}

	typ := reflect.TypeFor[T]()

	rows := readJsonTags(typ)
	vals := getRecordValues(r, rows)

	respJson, err := json.MarshalIndent(vals, "", "  ")
	if err != nil {
		return fail(fmt.Sprintf("marshal JSON from DB: %s", err.Error()))
	}

	fmt.Printf("DB RECORD: \n %s", string(respJson))

	return utils.Unmarshal[T](respJson)
}

func readJsonTags(typ reflect.Type) []string {
	rows := []string{}
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		jsonTag := field.Tag.Get("json")
		rows = append(rows, jsonTag)
	}
	return rows
}

func getRecordValues(r *db.Record, rows []string) map[string]interface{} {
	vals := make(map[string]interface{})
	for _, key := range rows {
		val, _ := r.Get(key)
		vals[key] = val
	}
	return vals
}

func mapToProperties[T any](params T) (map[string]interface{}, error) {
	marshaled, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	props, err := utils.Unmarshal[map[string]interface{}](marshaled)
	if err != nil {
		return nil, err
	}

	return props, nil
}